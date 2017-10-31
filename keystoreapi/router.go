package keystoreAPI

import "github.com/gorilla/mux"

//GetKeystoreRouter returns the API's router
func GetKeystoreRouter() *mux.Router {

	router := mux.NewRouter()

	router.Methods("GET").Path("/keys").HandlerFunc(getKeyList)
	router.Methods("POST").Path("/key").HandlerFunc(writeKey)
	router.Methods("GET").Path("/key/{key_id}").HandlerFunc(readKey)

	return router
}
