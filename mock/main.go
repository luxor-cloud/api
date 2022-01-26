package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "http listen address")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", getUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/servers", getServersHandler).Methods(http.MethodGet)
	r.HandleFunc("/servers/{id}", getServerHandler).Methods(http.MethodGet)
	r.HandleFunc("/servers/{id}/action", postActionHandler).Methods(http.MethodPost)
	r.HandleFunc("/servers/{id}/log", getLogsHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func getServersHandler(w http.ResponseWriter, r *http.Request) {

}

func getServerHandler(w http.ResponseWriter, r *http.Request) {

}

func postActionHandler(w http.ResponseWriter, r *http.Request) {

}

func getLogsHandler(w http.ResponseWriter, r *http.Request) {
	const op = "getLogsHandler"
	up := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		e(w, op, err)
		return
	}
	for {
		// TODO: send continous messages
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Hello World")); err != nil {
			log.Printf("%s: %v\n", op, err)
		}
	}
}

func e(w http.ResponseWriter, op string, err error) {
	log.Printf("%s: %v\n", op, err)
	// TODO: send error back
}
