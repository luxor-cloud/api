package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type action struct {
	Type string `json:"type"`
}

type apiErr struct {
	Message string `json:"msg"`
}

type logEntry struct {
	Ts   int64  `json:"ts"`
	Line string `json:"line"`
}

type serverLog struct {
	Entries []logEntry `json:"entries"`
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "http listen address")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", getUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/servers", getServersHandler).Methods(http.MethodGet)
	r.HandleFunc("/servers/{id}/action", postActionHandler).Methods(http.MethodPost)
	r.HandleFunc("/servers/{id}/log", getLogsHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	const op = "getUserHandler"
	id, ok := mux.Vars(r)["id"]
	if !ok {
		e(w, op, errors.New("id is missing"), http.StatusBadRequest)
		return
	}
	user, ok := users[id]
	if !ok {
		e(w, op, errors.New("no user found"), http.StatusNotFound)
		return
	}
	sendResp(w, user)
}

func getServersHandler(w http.ResponseWriter, r *http.Request) {
	const op = "getServersHandler"
	id, ok := mux.Vars(r)["id"]
	if !ok {
		e(w, op, errors.New("id is missing"), http.StatusBadRequest)
		return
	}
	server, ok := serversByUser[id]
	if !ok {
		e(w, op, errors.New("no user found"), http.StatusNotFound)
		return
	}
	sendResp(w, server)
}

func postActionHandler(w http.ResponseWriter, r *http.Request) {
	const op = "postActionHandler"
	all := make(map[string]server, 0)
	for _, v := range serversByUser {
		for _, server := range v {
			all[server.ID] = server
		}
	}
	id, ok := mux.Vars(r)["id"]
	if !ok {
		e(w, op, errors.New("id is missing"), http.StatusBadRequest)
		return
	}
	_, ok = all[id]
	if !ok {
		e(w, op, errors.New("no user found"), http.StatusNotFound)
		return
	}
	// TODO: simulate action stuff
}

func getLogsHandler(w http.ResponseWriter, r *http.Request) {
	const op = "getLogsHandler"
	up := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		e(w, op, err, http.StatusInternalServerError)
		return
	}
	for {
		// max 10 lines, min 1 line
		lines := rand.Intn(10-1) + 1
		entries := make([]logEntry, 0)
		ts := time.Now().UnixNano()

		for i := 0; i < lines; i++ {
			entries = append(entries, logEntry{
				Ts:   ts,
				Line: lorem.Sentence(1, 20),
			})
		}

		data, err := json.Marshal(serverLog{Entries: entries})
		if err != nil {
			log.Printf("%s: %v\n", op, err)
			return
		}

		wait := rand.Intn(3-1) + 1
		time.Sleep(time.Duration(wait) * time.Second)

		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Printf("%s: %v\n", op, err)
		}
	}
}

func sendResp(w http.ResponseWriter, body interface{}) {
	const op = "sendResp"
	j, err := json.Marshal(body)
	if err != nil {
		e(w, op, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(j); err != nil {
		log.Printf("%s: %v\n", op, err)
	}
}

func e(w http.ResponseWriter, op string, err error, code int) {
	log.Printf("%s: %v\n", op, err)
	// TODO: send error back
}
