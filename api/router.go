package api

import "github.com/gorilla/mux"

// Router returns the api's router
func Router() *mux.Router {
	router := mux.NewRouter()

	router.Methods("POST").Path("/key").HandlerFunc(writeKey)
	router.Methods("GET").Path("/key/{key_id}").HandlerFunc(readKey)
	router.Methods("GET").Path("/keys").HandlerFunc(readKeys)

	return router
}
