package event

import "sync"

type Stream struct {
	url      string
	pollOpts []PollOption
	poller   *Poller

	lastQueueEvent       *QueueEvent
	lastNowPlayingEvent  *NowPlayingEvent
	lastSessionDataEvent *SessionDataEvent

	mu             sync.Mutex
	nextObserverID int
	observers      map[int]*StreamObserver
}

func NewStream(url string, pollOpts ...PollOption) *Stream {
	return &Stream{
		url:       url,
		pollOpts:  pollOpts,
		observers: make(map[int]*StreamObserver),
	}
}

func (s *Stream) monitor() {
	for {
		select {
		case event, ok := <-s.poller.QueueEvents():
			if !ok {
				return
			}

			s.mu.Lock()
			s.lastQueueEvent = &event
			for _, obs := range s.observers {
				select {
				case obs.queueEvents <- event:
				default:
				}
			}
			s.mu.Unlock()
		case event, ok := <-s.poller.NowPlayingEvents():
			if !ok {
				return
			}

			s.mu.Lock()
			s.lastNowPlayingEvent = &event
			for _, obs := range s.observers {
				select {
				case obs.nowPlayingEvents <- event:
				default:
				}
			}
			s.mu.Unlock()
		case event, ok := <-s.poller.SessionDataEvents():
			if !ok {
				return
			}

			s.mu.Lock()
			s.lastSessionDataEvent = &event
			for _, obs := range s.observers {
				select {
				case obs.sessionDataEvents <- event:
				default:
				}
			}
			s.mu.Unlock()
		}
	}
}

func (s *Stream) NewObserver() (*StreamObserver, error) {
	obs := &StreamObserver{
		stream:            s,
		queueEvents:       make(chan QueueEvent, 1),
		nowPlayingEvents:  make(chan NowPlayingEvent, 1),
		sessionDataEvents: make(chan SessionDataEvent, 1),
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Do we have a poller? If not, we need to create one
	if s.poller == nil {
		poller, err := NewPoller(s.url, s.pollOpts...)
		if err != nil {
			return nil, err
		}

		s.poller = poller

		go s.monitor()
	}

	// Link observers to the current stream.
	obs.observerID = s.nextObserverID
	s.observers[obs.observerID] = obs
	s.nextObserverID++

	// If there is any data, prepopulate the channel
	if s.lastQueueEvent != nil {
		obs.queueEvents <- *s.lastQueueEvent
	}
	if s.lastNowPlayingEvent != nil {
		obs.nowPlayingEvents <- *s.lastNowPlayingEvent
	}
	if s.lastQueueEvent != nil {
		obs.sessionDataEvents <- *s.lastSessionDataEvent
	}

	return obs, nil
}

type StreamObserver struct {
	stream *Stream

	observerID        int
	queueEvents       chan QueueEvent
	nowPlayingEvents  chan NowPlayingEvent
	sessionDataEvents chan SessionDataEvent
}

func (s *StreamObserver) QueueEvents() <-chan QueueEvent             { return s.queueEvents }
func (s *StreamObserver) NowPlayingEvents() <-chan NowPlayingEvent   { return s.nowPlayingEvents }
func (s *StreamObserver) SessionDataEvents() <-chan SessionDataEvent { return s.sessionDataEvents }

func (s *StreamObserver) Close() error {
	s.stream.mu.Lock()
	defer s.stream.mu.Unlock()

	// Are we the last one to leave? If so, shutdown the poller
	if len(s.stream.observers) == 1 {
		s.stream.poller.Close()
		s.stream.poller = nil
	}

	delete(s.stream.observers, s.observerID)
	close(s.queueEvents)
	close(s.nowPlayingEvents)
	close(s.sessionDataEvents)

	return nil
}
