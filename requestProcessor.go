package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/glog"
)

const (
	transactionServerUrl = "http://localhost:9090/api/"
	addPath              = "add"
	authenticatePath     = "authenticate"
)

// FUNCTION TO SENT A HTTP REQUEST
func sendHttpRequest(requestBody map[string]interface{}, path string) *http.Response {
	bytesRepresentation, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(transactionServerUrl+path, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

type RequestQueue struct {
}

var UserRequestsMap = make(map[string]RequestQueue)

// Authenticate User
func authenticateUser(userId string) {
	glog.Info("############################### INFO: Executing authenticate for: ", userId)

	requestBody := map[string]interface{}{
		"UserId": userId}

	var result = sendHttpRequest(requestBody, authenticatePath)

	if true {
		glog.Info("############################### SUCCESS: Authentication Successful!", result)
	} else {
		glog.Info("############################### ERROR: Authentication Unsuccessful!")
	}
}

func addFunds(userId string, priceDollars float64, priceCents float64, commandNumber int) {
	glog.Info("\n\n############################### INFO: Creating Add HTTP Request...")
	requestBody := map[string]interface{}{
		"UserId":        userId,
		"PriceDollars":  priceDollars,
		"CommandNumber": commandNumber}

	sendHttpRequest(requestBody, addPath)

}
