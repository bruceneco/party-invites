// Package handlers Party Invites API.
//
// This application will manage invites for a party.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
//     Contact: Bruce Neco<bruce.neco3@outlook.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handler

import "github.com/gorilla/mux"

type Handler interface {
	Register(sm *mux.Router)
}
