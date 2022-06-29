package handler

import (
	"github.com/bruceneco/party-invites/service"
	"github.com/gorilla/mux"
	"net/http"
)

type guestHandler struct {
	sp *service.ServiceProvider
}

func NewGuestHandler(sp *service.ServiceProvider) Handler {
	return &guestHandler{sp: sp}
}
func (gh *guestHandler) Register(sm *mux.Router) {
	sm.HandleFunc("/", gh.getGuests).Methods(http.MethodGet)
	sm.HandleFunc("/{id:\\d+}", gh.getGuestById)
}
