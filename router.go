package main

import "github.com/gorilla/mux"

func getRouter() *mux.Router {

	router := mux.NewRouter()

	router.Methods("GET").Path("/keys").HandlerFunc(getKeyList)
	router.Methods("POST").Path("/key").HandlerFunc(writeKey)
	router.Methods("GET").Path("/key/{key_id}").HandlerFunc(readKey)

	return router
}
