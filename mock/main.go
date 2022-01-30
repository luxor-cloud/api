package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var jwtSigningKey = []byte("test")

type userClaims struct {
	jwt.StandardClaims
	ID string `json:"userId"`
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type server struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IP         string `json:"ip"`
	Region     string `json:"region"`
	Datacenter string `json:"datacenter"`
	Type       string `json:"type"`
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

type createServerReq struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "http listen address")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/me", getUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/servers", getServersHandler).Methods(http.MethodGet)
	r.HandleFunc("/servers", putServersHandler).Methods(http.MethodPut)
	r.HandleFunc("/servers/{id}/action", postActionHandler).Methods(http.MethodPost)
	r.HandleFunc("/servers/{id}/log", getLogsHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	const op = "getUserHandler"
	id, err := userIDFromHeader(r)
	if err != nil {
		e(w, op, err, http.StatusBadRequest)
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
	id, err := userIDFromHeader(r)
	if err != nil {
		e(w, op, err, http.StatusBadRequest)
		return
	}
	server, ok := serversByUser[id]
	if !ok {
		e(w, op, errors.New("no user found"), http.StatusNotFound)
		return
	}
	sendResp(w, server)
}

func putServersHandler(w http.ResponseWriter, r *http.Request) {
	const op = "putServersHandler"
	id, err := userIDFromHeader(r)
	if err != nil {
		e(w, op, err, http.StatusBadRequest)
		return
	}

	tmp, err := readReq(r)
	if err != nil {
		e(w, op, err, http.StatusInternalServerError)
	}

	req := tmp.(createServerReq)

	serverID, err := uuid.NewRandom()
	if err != nil {
		e(w, op, err, http.StatusBadRequest)
		return
	}

	created := server{
		ID:         serverID.String(),
		Name:       req.Name,
		IP:         "45.30.234",
		Region:     req.Region,
		Datacenter: "my-dc15",
		Type:       req.Type,
	}

	serversByUser[id] = append(serversByUser[id], created)

	// simulate some response time
	wait := rand.Intn(5-1) + 1
	time.Sleep(time.Duration(wait) * time.Second)

	sendResp(w, created)
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

func userIDFromHeader(r *http.Request) (string, error) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return "", errors.New("no authorization header set")
	}

	// format: bearer <token>
	parts := strings.Split(tokenStr, " ")
	if len(parts) < 2 {
		return "", errors.New("authorization header contains invalid value")
	}

	token, err := jwt.ParseWithClaims(parts[1], &userClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*userClaims)
	if !ok || !token.Valid {
		return "", errors.New("token or claims not valid")
	}
	return claims.ID, nil
}

func readReq(r *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return v, nil
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
	payload := fmt.Sprintf(`{"msg": "%v"}`, err)
	w.WriteHeader(code)
	if _, err := w.Write([]byte(payload)); err != nil {
		log.Printf("%s: %v\n", op, err)
	}
}
