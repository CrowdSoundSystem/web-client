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

func eventStreamHandler(ws *websocket.Conn) {
	defer ws.Close()

	obs, err := eventStream.NewObserver()
	if err != nil {
		b, _ := json.Marshal(event.EventData{Error: err})
		io.Copy(ws, bytes.NewReader(b))
		return
	}
	defer obs.Close()

	for event := range obs.Events() {
		serialized, _ := json.Marshal(event)
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
