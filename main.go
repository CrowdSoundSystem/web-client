package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/crowdsoundsystem/web-client/pkg/event"
	"golang.org/x/net/websocket"
)

var (
	eventStream *event.Stream
)

type EventData struct {
	EventType string      `json:"eventType,omitempty"`
	Error     error       `json:"error,omitempty"`
	Event     interface{} `json:"event,omitempty"`
}

func eventStreamHandler(ws *websocket.Conn) {
	defer ws.Close()

	obs, err := eventStream.NewObserver()
	if err != nil {
		b, _ := json.Marshal(EventData{Error: err})
		io.Copy(ws, bytes.NewReader(b))
		return
	}
	defer obs.Close()

	for {
		var serialized []byte
		select {
		case event, ok := <-obs.QueueEvents():
			if !ok {
				return
			}
			serialized, _ = json.Marshal(EventData{
				EventType: "queue",
				Event:     event,
			})
		case event, ok := <-obs.NowPlayingEvents():
			if !ok {
				return
			}
			serialized, _ = json.Marshal(EventData{
				EventType: "now_playing",
				Event:     event,
			})
		case event, ok := <-obs.SessionDataEvents():
			if !ok {
				return
			}
			serialized, _ = json.Marshal(EventData{
				EventType: "session_data",
				Event:     event,
			})
		}

		io.Copy(ws, bytes.NewReader(serialized))
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	eventStream = event.NewStream("cs.ephyra.io:50051")

	http.Handle("/event_stream", websocket.Handler(eventStreamHandler))
	http.Handle("/", http.FileServer(http.Dir("site/")))

	panic(http.ListenAndServe(":8080", nil))
}
