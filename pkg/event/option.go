package event

import "time"

type pollerConfig struct {
	queueInterval       time.Duration
	nowPlayingInterval  time.Duration
	sessionDataInterval time.Duration
}

type PollOption func(*pollerConfig)

func WithQueueInterval(i time.Duration) PollOption {
	return func(p *pollerConfig) {
		p.queueInterval = i
	}
}

func WithNowPlayingInterval(i time.Duration) PollOption {
	return func(p *pollerConfig) {
		p.nowPlayingInterval = i
	}
}

func WithSessionDataInterval(i time.Duration) PollOption {
	return func(p *pollerConfig) {
		p.sessionDataInterval = i
	}
}
