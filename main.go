package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adrianosela/Keystore/keystoreapi"
)

func main() {
	//check the keys directory is in place
	err := checkPreconditions()
	if err != nil {
		log.Fatalf("Could not read keys directory. %s", err)
	}

	router := keystoreAPI.GetKeystoreRouter()

	log.Println("[INFO] Listening on http://localhost:80")

	err = http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func checkPreconditions() error {
	_, err := ioutil.ReadDir("./keys")
	return err
}
