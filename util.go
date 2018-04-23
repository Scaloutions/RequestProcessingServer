package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/glog"
)

const (
	transactionServerUrl = "http://transactionserver:9090/api/" //"http://localhost:9090/api/"
	// transactionServerUrl = "http://localhost:9090/api/"
	ADD              = "add"
	QUOTE            = "quote"
	BUY              = "buy"
	COMMIT_BUY       = "commit_buy"
	CANCEL_BUY       = "cancel_buy"
	SELL             = "sell"
	COMMIT_SELL      = "commit_sell"
	CANCEL_SELL      = "cancel_sell"
	SET_BUY_AMOUNT   = "set_buy_amount"
	CANCEL_SET_BUY   = "cancel_set_buy"
	SET_BUY_TRIGGER  = "set_buy_trigger"
	SET_SELL_AMOUNT  = "set_sell_amount"
	SET_SELL_TRIGGER = "set_sell_trigger"
	CANCEL_SET_SELL  = "cancel_set_sell"
	DISPLAY_SUMMARY  = "display_summary"
	DUMPLOG          = "dumplog"
	authenticatePath = "authenticate"
)

// FUNCTION TO SENT A HTTP REQUEST
func sendHttpRequest(requestBody map[string]interface{}, path string) *http.Response {
	bytesRepresentation, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	// resp, err := http.Post(transactionServerUrl+path, "application/json", bytes.NewBuffer(bytesRepresentation))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	req, reqErr := http.NewRequest("POST", transactionServerUrl+path, bytes.NewBuffer(bytesRepresentation))
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		glog.Error("Error sending a POST request to Transaction server", err)
	}

	req.Close = true

	return resp
}
