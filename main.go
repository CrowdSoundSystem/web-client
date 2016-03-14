package main

func main() {
	// TODO: Create websocket event model
	// TODO: Create API's for crowdsound management. This requires crowdsound admin service.
	// TODO: Serve index and settings page (can just use caching?).

	// The websocket model is simply an event stream, whereby multiple polling goroutines
	// of the crowdsound service report any changes to clients. This approach is very similar
	// to the logstreamer implementation (with the exception of a single last known state).
	//
	// Alternatively, the stream can powered by goroutines dedicated to a given client. Number
	// of web clients in general won't exceed single (at worst, double) digits. However, it would
	// still be nice have some kind of efficiency (maybe later).

	// The API for crowdsound management is simple, as it is just translating HTTP/1.1 requests
	// into a gRPC request. A global client would probably suffice, to keep connections low.

	// Serving pages are easy, just reply with the bytess.
}
