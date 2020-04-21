package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router returns the keystore api's router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/key").HandlerFunc(writeKey)
	r.Methods(http.MethodGet).Path("/key/{key_id}").HandlerFunc(readKey)
	r.Methods(http.MethodGet).Path("/keys").HandlerFunc(readKeys)
	return r
}
