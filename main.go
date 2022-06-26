package main

import (
	"fmt"
	"net/http"

	"github.com/bruceneco/party-invites/handler"
	"github.com/bruceneco/party-invites/repository"
	"github.com/bruceneco/party-invites/utils"
	"github.com/gorilla/mux"
)

func main() {
	utils.InitLogger()

	cfg, err := utils.LoadConfig()
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}

	fp := "./db.sqlite"
	repository.NewDB(fp)

	sm := mux.NewRouter()

	handler.NewGuestHandler(sm.PathPrefix("/guest").Subrouter())

	host := cfg.Server.Host
	port := cfg.Server.Port
	utils.InfoLog.Printf("Server started at %s:%s\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
}
