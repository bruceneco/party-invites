package main

import (
	"fmt"
	"net/http"

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
	dao := repository.NewDAO()

	sm := mux.NewRouter()

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gq := dao.NewGuestQuery()
		id, err := gq.NewGuest()
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			utils.ErrorLog.Print(err)
			return
		}
		w.Write([]byte(string(id)))
	})

	host := cfg.Server.Host
	port := cfg.Server.Port
	utils.InfoLog.Printf("Server started at %s:%s\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), sm)
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
}
