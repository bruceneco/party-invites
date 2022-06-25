package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bruceneco/party-invites/utils"
	"github.com/gorilla/mux"
)

func main() {
	l := log.Logger{}
	utils.InitLogger(&l)
	cfg, err := utils.LoadConfig()
	if err != nil {
		l.Fatalln(err)
	}

	host := cfg.Server.Host
	port := cfg.Server.Port

	sm := mux.NewRouter()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gopher!"))
	})
	log.Printf("[INFO] Server started at %s:%s\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		l.Fatalln(err)
	}
}
