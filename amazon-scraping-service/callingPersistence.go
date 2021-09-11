package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func callPersistence(Response *responsePayload) string {
	jsonReq, err := json.Marshal(Response)
	resp, err := http.Post("http://persister:8081/persistProduct", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString
}
