package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func eventStreamHandler(ws *websocket.Conn) {
	// The websocket model is simply an event stream, whereby multiple polling goroutines
	// of the crowdsound service report any changes to clients. This approach is very similar
	// to the logstreamer implementation (with the exception of a single last known state).
	//
	// Alternatively, the stream can powered by goroutines dedicated to a given client. Number
	// of web clients in general won't exceed single (at worst, double) digits. However, it would
	// still be nice have some kind of efficiency (maybe later).
	//
	// TODO: Select one, implement it, and hook it up here.
}

func main() {
	// TODO: Create websocket event model
	// TODO: Create API's for crowdsound management. This requires crowdsound admin service.
	// TODO: Serve index and settings page (can just use caching?).

	// The API for crowdsound management is simple, as it is just translating HTTP/1.1 requests
	// into a gRPC request. A global client would probably suffice, to keep connections low.
	http.Handle("/event_stream", websocket.Handler(eventStreamHandler))
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("site/"))))
}
