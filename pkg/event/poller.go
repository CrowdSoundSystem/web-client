package event

import (
	"io"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/crowdsoundsystem/go-client/pkg/crowdsound"
)

type Poller struct {
	config pollerConfig
	conn   *grpc.ClientConn
	client crowdsound.CrowdSoundClient

	shutdown  chan struct{}
	eventChan chan EventData
}

func NewPoller(url string, options ...PollOption) (*Poller, error) {
	config := pollerConfig{
		queueInterval:       2 * time.Second,
		nowPlayingInterval:  2 * time.Second,
		sessionDataInterval: 5 * time.Second,
	}

	// Apply caller specified options
	for _, opt := range options {
		opt(&config)
	}

	// Create gRPC connection to crowdsound.
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	poller := &Poller{
		config:    config,
		conn:      conn,
		client:    crowdsound.NewCrowdSoundClient(conn),
		shutdown:  make(chan struct{}),
		eventChan: make(chan EventData),
	}

	go poller.pollQueue()
	go poller.pollNowPlaying()
	go poller.pollSessionData()

	return poller, nil
}

func (p *Poller) Close() error {
	close(p.shutdown)
	close(p.eventChan)

	return p.conn.Close()
}

func (p *Poller) pollQueue() {
	for {
		select {
		case <-p.shutdown:
			return
		case <-time.After(p.config.queueInterval):
			stream, err := p.client.GetQueue(context.Background(), &crowdsound.GetQueueRequest{})
			if err != nil {
				log.Println("Unable to retreive now playing:", err)
				continue
			}

			event := &QueueEvent{}

			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Println("Unexpected error streaming getQueue:", err)
					break
				}

				song := Song{
					Name:   resp.Name,
					Artist: resp.Artist,
					Genre:  resp.Genre,
				}

				if resp.IsBuffered {
					event.Buffered = append(event.Buffered, song)
				} else {
					event.Queued = append(event.Queued, song)
				}
			}

			p.eventChan <- EventData{
				EventType: "queue",
				Event:     event,
			}
		}
	}
}

func (p *Poller) pollNowPlaying() {
	for {
		select {
		case <-p.shutdown:
			return
		case <-time.After(p.config.nowPlayingInterval):
			resp, err := p.client.GetPlaying(context.Background(), &crowdsound.GetPlayingRequest{})
			if err != nil {
				log.Println("Unable to retreive now playing:", err)
				continue
			}

			p.eventChan <- EventData{
				EventType: "now_playing",
				Event: &NowPlayingEvent{
					Song: Song{Name: resp.Name, Artist: resp.Artist, Genre: resp.Genre},
				},
			}
		}
	}
}

func (p *Poller) pollSessionData() {
	for {
		select {
		case <-p.shutdown:
			return
		case <-time.After(p.config.sessionDataInterval):
			resp, err := p.client.GetSessionData(context.Background(), &crowdsound.GetSessionDataRequest{})
			if err != nil {
				log.Println("Unable to retrieve session data:", err)
				continue
			}

			p.eventChan <- EventData{
				EventType: "sesssion_data",
				Event: &SessionDataEvent{
					SessionName: resp.SessionName,
					Users:       int(resp.NumUsers),
				},
			}
		}
	}
}
