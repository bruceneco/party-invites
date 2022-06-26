package main

import (
	"fmt"
	"net/http"

	"github.com/bruceneco/party-invites/utils"
	"github.com/gorilla/mux"
)

var cfg *utils.Config
var sm http.Handler

func main() {
	stepLog()
	stepEnv()
	stepServer()
	stepListenAndServe()
}

func stepLog() {
	utils.InitLogger()
}

func stepEnv() {
	var err error
	cfg, err = utils.LoadConfig()
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
}

func stepServer() {
	sm := mux.NewRouter()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gopher!"))
	})
}

func stepListenAndServe() {
	host := cfg.Server.Host
	port := cfg.Server.Port
	utils.InfoLog.Printf("Server started at %s:%s\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
}
