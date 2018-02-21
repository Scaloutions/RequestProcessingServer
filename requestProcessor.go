package main

import (
	"log"
	"github.com/golang/glog"
	"net/http"
	"bytes"
	"encoding/json"
)

const (
	transactionServerUrl = "http://localhost:9090/api/"
	addPath = "add"
	authenticatePath = "authenticate"
)

type RequestQueue struct {
	
}


var UserRequestsMap = make(map[string]RequestQueue)

func sendHttpRequest(requestBody map[string]interface{}, path string) map[string]interface{}  {
	bytesRepresentation, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(transactionServerUrl + path, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}


	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)


	glog.Info("$$$$$$$$$$$$$$$$$$$ RETURNED ")
	log.Println("\n\n\n\n\n\n $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$", result, len(result))
	glog.Info("$$$$$$$$$$$$$$$$$$$ DONE ")
	
	return result
}


// Authenticate User 
func authenticateUser (userId string) {
	glog.Info("############################### INFO: Executing authenticate for: ", userId)
	
	requestBody := map[string]interface{} { 
		"UserId": userId }

	var result = sendHttpRequest(requestBody, authenticatePath)

  if (true) {
    glog.Info("############################### SUCCESS: Authentication Successful!", result)
  } else {
    glog.Info("############################### ERROR: Authentication Unsuccessful!")
  }
}

func addFunds(userId string, priceDollars float64, priceCents float64, commandNumber int) {
	glog.Info("\n\n############################### INFO: Creating Add HTTP Request...")
	requestBody := map[string]interface{} {
		"UserId": userId,
		"PriceDollars": priceDollars,
		"CommandNumber": commandNumber }
	
	sendHttpRequest(requestBody, addPath)

}

