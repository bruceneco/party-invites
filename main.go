package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bruceneco/party-invites/utils"
	"github.com/gorilla/mux"
)

func main() {

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	host := cfg.Server.Host
	port := cfg.Server.Port

	sm := mux.NewRouter()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gopher!"))
	})
	log.Printf("[INFO] Server started at port %s:%s\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		log.Fatalln(err)
	}

}
