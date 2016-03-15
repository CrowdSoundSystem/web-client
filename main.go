package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/crowdsoundsystem/web-client/pkg/crowdsound"
	"github.com/crowdsoundsystem/web-client/pkg/event"
	"github.com/crowdsoundsystem/web-client/pkg/settings"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
)

var (
	eventStream    *event.Stream
	remoteSettings *settings.Settings

	endpoint = flag.String("endpoint", "localhost:50051", "Crowdsound endpoint")
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

func skipHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := crowdsound.NewAdminClient(conn)
	_, err = client.Skip(context.Background(), &crowdsound.SkipRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func settingHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		s, err := remoteSettings.Get(false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		serialized, _ := json.Marshal(s)
		io.Copy(w, bytes.NewReader(serialized))
	} else if req.Method == "POST" {
		key := req.FormValue("key")
		valType := req.FormValue("type")
		val := req.FormValue("value")

		if key == "" || valType == "" || val == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		var err error

		switch valType {
		case "bool":
			boolVal, err := strconv.ParseBool(val)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = remoteSettings.SetBool(key, boolVal)
			if err != nil {
				log.Println("err: error")
			}
			break
		case "int":
			intVal, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = remoteSettings.SetInt(key, int(intVal))
			break
		case "float":
			floatVal, err := strconv.ParseFloat(val, 32)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = remoteSettings.SetFloat(key, float32(floatVal))
			break
		case "string":
			err = remoteSettings.SetString(key, val)
			break
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := crowdsound.NewAdminClient(conn)
	resp, err := client.GetVersionInfo(context.Background(), &crowdsound.GetVersionInfoRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(resp)
	io.Copy(w, bytes.NewReader(serialized))
}

func dbStatsHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := crowdsound.NewAdminClient(conn)
	resp, err := client.GetDBStats(context.Background(), &crowdsound.GetDBStatsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(resp)
	io.Copy(w, bytes.NewReader(serialized))
}

func main() {
	flag.Parse()

	var err error
	eventStream = event.NewStream(*endpoint)
	remoteSettings, err = settings.NewSettings(*endpoint, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/admin/skip", skipHandler)
	http.HandleFunc("/admin/setting", settingHandler)
	http.HandleFunc("/admin/version", versionHandler)
	http.HandleFunc("/admin/db_stats", dbStatsHandler)
	http.Handle("/event_stream", websocket.Handler(eventStreamHandler))
	http.Handle("/", http.FileServer(http.Dir("site/")))

	panic(http.ListenAndServe(":8080", nil))
}
