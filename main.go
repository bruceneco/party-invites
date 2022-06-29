package main

import (
	"fmt"
	"github.com/bruceneco/party-invites/service"
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

	// Repo
	fp := "./db.sqlite"
	repository.NewDB(fp)
	dao := repository.NewDAO()

	// Services
	guestService := service.NewGuestService(dao)
	sp := service.NewServiceProvider(guestService)

	// Router
	sm := mux.NewRouter().StrictSlash(true)

	// Handlers
	guestHandler := handler.NewGuestHandler(sp)
	guestHandler.Register(sm.PathPrefix("/guest").Subrouter())

	// Server
	host := cfg.Server.Host
	port := cfg.Server.Port
	utils.InfoLog.Printf("Server started at %s:%s\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
}
