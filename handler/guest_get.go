package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bruceneco/party-invites/dto"
	"github.com/bruceneco/party-invites/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (gh *guestHandler) getGuests(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	guests, err := gh.sp.GuestService.GetGuests()
	if err != nil {
		utils.ErrorLog.Print(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}
	e := json.NewEncoder(w)
	err = e.Encode(guests)
	if err != nil {
		utils.ErrorLog.Print(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}
}

func (gh *guestHandler) getGuestById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLog.Print(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}
	var g *dto.Guest
	g, err = gh.sp.GuestService.GetGuest(id)
	if err != nil {
		utils.ErrorLog.Print(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}
	e := json.NewEncoder(w)
	err = e.Encode(g)
	if err != nil {
		utils.ErrorLog.Print(err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}
}
