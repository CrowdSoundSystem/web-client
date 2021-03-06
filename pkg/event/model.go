package event

type EventData struct {
	EventType string      `json:"eventType,omitempty"`
	Error     error       `json:"error,omitempty"`
	Event     interface{} `json:"event,omitempty"`
}

type Song struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
}

type QueueEvent struct {
	Buffered []Song `json:"buffered"`
	Queued   []Song `json:"queued"`
}

type NowPlayingEvent struct {
	Song Song `json:"song"`
}

type SessionDataEvent struct {
	SessionName string `json:"sessionName"`
	Users       int    `json:"users"`
}

type TrendingArtistsEvent struct {
	Artists []string `json:"artists"`
}

type SkipStatusEvent struct {
	VotesToSkip int     `json:"vote_to_skip"`
	TotalUsers  int     `json:"total_users"`
	Threshold   float32 `json:"threshold"`
}
