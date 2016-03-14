package event

import "sync"

type Stream struct {
	url      string
	pollOpts []PollOption
	poller   *Poller

	lastEvents map[string]EventData

	mu             sync.Mutex
	nextObserverID int
	observers      map[int]*StreamObserver
}

func NewStream(url string, pollOpts ...PollOption) *Stream {
	return &Stream{
		url:        url,
		pollOpts:   pollOpts,
		lastEvents: make(map[string]EventData),
		observers:  make(map[int]*StreamObserver),
	}
}

func (s *Stream) monitor() {
	for {
		select {
		case event, ok := <-s.poller.eventChan:
			if !ok {
				return
			}

			s.mu.Lock()

			s.lastEvents[event.EventType] = event
			for _, obs := range s.observers {
				select {
				case obs.eventsChan <- event:
				default:
				}
			}
			s.mu.Unlock()
		}
	}
}

func (s *Stream) NewObserver() (*StreamObserver, error) {
	obs := &StreamObserver{
		stream: s,
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

	obs.eventsChan = make(chan EventData, len(s.lastEvents))
	for _, event := range s.lastEvents {
		obs.eventsChan <- event
	}

	return obs, nil
}

type StreamObserver struct {
	stream *Stream

	observerID int
	eventsChan chan EventData
}

func (s *StreamObserver) Events() <-chan EventData { return s.eventsChan }

func (s *StreamObserver) Close() error {
	s.stream.mu.Lock()
	defer s.stream.mu.Unlock()

	// Are we the last one to leave? If so, shutdown the poller
	if len(s.stream.observers) == 1 {
		s.stream.poller.Close()
		s.stream.poller = nil
	}

	delete(s.stream.observers, s.observerID)
	close(s.eventsChan)

	return nil
}
