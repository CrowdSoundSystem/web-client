package event

import (
	"io"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/crowdsoundsystem/web-client/pkg/crowdsound"
)

type Poller struct {
	config           pollerConfig
	conn             *grpc.ClientConn
	crowdsoundClient crowdsound.CrowdSoundClient
	adminClient      crowdsound.AdminClient

	shutdown  chan struct{}
	eventChan chan EventData
}

func NewPoller(url string, options ...PollOption) (*Poller, error) {
	config := pollerConfig{
		queueInterval:           2 * time.Second,
		nowPlayingInterval:      2 * time.Second,
		sessionDataInterval:     2 * time.Second,
		trendingArtistsInterval: 2 * time.Second,
		skipStatusInterval:      1 * time.Second,
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
		config:           config,
		conn:             conn,
		crowdsoundClient: crowdsound.NewCrowdSoundClient(conn),
		adminClient:      crowdsound.NewAdminClient(conn),
		shutdown:         make(chan struct{}),
		eventChan:        make(chan EventData),
	}

	go poller.pollQueue()
	go poller.pollNowPlaying()
	go poller.pollSessionData()
	go poller.pollTrendingArtists()
	go poller.pollSkipStatus()

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
			stream, err := p.crowdsoundClient.GetQueue(context.Background(), &crowdsound.GetQueueRequest{})
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
			resp, err := p.crowdsoundClient.GetPlaying(context.Background(), &crowdsound.GetPlayingRequest{})
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
			resp, err := p.crowdsoundClient.GetSessionData(context.Background(), &crowdsound.GetSessionDataRequest{})
			if err != nil {
				log.Println("Unable to retrieve session data:", err)
				continue
			}

			p.eventChan <- EventData{
				EventType: "session_data",
				Event: &SessionDataEvent{
					SessionName: resp.SessionName,
					Users:       int(resp.NumUsers),
				},
			}
		}
	}
}

func (p *Poller) pollTrendingArtists() {
	for {
		select {
		case <-p.shutdown:
			return
		case <-time.After(p.config.trendingArtistsInterval):
			stream, err := p.crowdsoundClient.ListTrendingArtists(context.Background(), &crowdsound.ListTrendingArtistsRequest{})
			if err != nil {
				log.Println("Unable to retrieve trending artists:", err)
				continue
			}

			event := &TrendingArtistsEvent{}

			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Println("Unexpected error streaming trendingArtist:", err)
					break
				}

				event.Artists = append(event.Artists, resp.Name)
			}

			p.eventChan <- EventData{
				EventType: "trending_artists",
				Event:     event,
			}
		}
	}
}

func (p *Poller) pollSkipStatus() {
	for {
		select {
		case <-p.shutdown:
			return
		case <-time.After(p.config.skipStatusInterval):
			resp, err := p.adminClient.SkipStatus(context.Background(), &crowdsound.SkipStatusRequest{})
			if err != nil {
				log.Println("Unable to retrieve skip status:, err")
				continue
			}

			p.eventChan <- EventData{
				EventType: "skip_status",
				Event: &SkipStatusEvent{
					VotesToSkip: int(resp.VotesToSkip),
					TotalUsers:  int(resp.TotalUsers),
					Threshold:   resp.Threshold,
				},
			}
		}
	}
}
