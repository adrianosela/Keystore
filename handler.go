package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getKeyList(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./keys")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not get keys")
		return
	}

	jsonArray := struct {
		KeyIDList []string `json:"keys"`
	}{
		KeyIDList: []string{},
	}

	for _, f := range files {
		jsonArray.KeyIDList = append(jsonArray.KeyIDList, f.Name())
	}

	bytes, err := json.Marshal(jsonArray)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not marshall keys")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bytes))
	return
}

func writeKey(w http.ResponseWriter, r *http.Request) {
	keyID := "mock"

	f, err := os.Create("./keys/" + keyID)
	defer f.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not create new file")
		return
	}
	//read bytes off request and write them to the file
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not read request")
		return
	}
	_, err = f.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not write to file")
		return
	}
}

func readKey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "mockkey")
	return
}
