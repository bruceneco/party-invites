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

// swagger:response getGuestsRes
type _getGuestsRes struct {
	//in:body
	Body []dto.Guest
}

// swagger:route GET /guests guests guestList
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: genericError
//       200: getGuestsRes
//       422: validationError
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

// swagger:response getGuestByIdRes
type _getGuestByIdRes struct {
	//in:body
	Body dto.Guest
}

// swagger:route GET /guests/{id} guests guestById
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: genericError
//       200: getGuestByIdRes
//       422: validationError
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
