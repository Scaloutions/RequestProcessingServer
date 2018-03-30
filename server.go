package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func usage() {
	fmt.Println("usage: example -logtostderr=true -stderrthreshold=[INFO|WARN|FATAL|ERROR] -log_dir=[string]\n")
	flag.PrintDefaults()
}

// User map
var f *os.File = nil

var UserIdRequestQueueMap = map[string]*UserTSRequestsDetails{}

var transactionHistory []string

type UserTSRequestsDetails struct {
	UserRequests     *RequestsQueue
	RunningGoRoutine bool
}

type Response struct {
	UserId        string
	PriceDollars  float64
	PriceCents    float64
	Command       string
	CommandNumber int
	Stock         string
}

func handleHttpRequestResponse(requestBody map[string]interface{}, path string, RequestsQueue *RequestsQueue, CommandNumber int) {
	glog.Info("################################################### LOOKING AT: ", path, CommandNumber)
	var httpResponse = sendHttpRequest(requestBody, path)
	var httpJsonResponse map[string]interface{}
	json.NewDecoder(httpResponse.Body).Decode(&httpJsonResponse)

	if httpResponse.StatusCode == 200 {
		glog.Info("\n\n Request was suceessful. Response is:", httpJsonResponse)
	} else {
		glog.Info("\n OOPS :(")
	}

	// glog.Info("######### BEFORE REMOVING")
	// RequestsQueue.printQueue()
	RequestsQueue.Dequeue()
	// glog.Info("######### AFTER REMOVING")
	// RequestsQueue.printQueue()

}

func startProcessingUser(userId string, RequestsQueue *RequestsQueue) {
	// Indicate that a new routine has been spin up for user
	UserIdRequestQueueMap[userId].RunningGoRoutine = true
	glog.Info("\n\n\n\n\n\n\n\n\n\n ########## Inside a new routine for user: ", userId)
	currentReqNode := RequestsQueue.Head()

	if currentReqNode == nil {
		// glog.Info("!!!!!!!!! HEAD IS NULLL")
	}

	for currentReqNode != nil {
		i := currentReqNode.value
		currentRequest := i.(TSRequest)

		switch currentRequest.Command {
		case "add":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "add", RequestsQueue, currentRequest.CommandNumber)
			break

		case "quote":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "quote", RequestsQueue, currentRequest.CommandNumber)
			break

		case "buy":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "buy", RequestsQueue, currentRequest.CommandNumber)
			break

		case "commit_buy":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "commit_buy", RequestsQueue, currentRequest.CommandNumber)
			break

		case "cancel_buy":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "cancel_buy", RequestsQueue, currentRequest.CommandNumber)
			break

		case "sell":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "sell", RequestsQueue, currentRequest.CommandNumber)
			break

		case "commit_sell":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "commit_sell", RequestsQueue, currentRequest.CommandNumber)
			break

		case "cancel_sell":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "cancel_sell", RequestsQueue, currentRequest.CommandNumber)
			break

		case "set_buy_amount":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "set_buy_amount", RequestsQueue, currentRequest.CommandNumber)
			break

		case "cancel_set_buy":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "cancel_set_buy", RequestsQueue, currentRequest.CommandNumber)
			break

		case "set_buy_trigger":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "set_buy_trigger", RequestsQueue, currentRequest.CommandNumber)
			break

		case "set_sell_amount":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "set_sell_amount", RequestsQueue, currentRequest.CommandNumber)
			break

		case "set_sell_trigger":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "set_sell_trigger", RequestsQueue, currentRequest.CommandNumber)
			break

		case "cancel_set_sell":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "cancel_set_sell", RequestsQueue, currentRequest.CommandNumber)
			break

		case "display_summary":
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, "display_summary", RequestsQueue, currentRequest.CommandNumber)
			break

		}
		currentReqNode = currentReqNode.next
	}

	glog.Info("EXITING ROUTINE FOR USER ", userId, " .BYEEEE")
	UserIdRequestQueueMap[userId].RunningGoRoutine = false

}

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.RequestURI)
// 		next.ServeHTTP(w, r)
// 	})
// }

//Parse request and return Response Object
// Get the user request -- done
// Add to user map -- done
// If there is a go routine thats servicing a user, do nothing
// If there is no go routine, start one up
// In the go routine, create the hhtp request and fire
func parseRequest(w http.ResponseWriter, r *http.Request) {
	msg := Response{}
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}

	var incomingReq = TSRequest{
		UserId:        msg.UserId,
		PriceDollars:  msg.PriceDollars,
		PriceCents:    msg.PriceCents,
		Command:       msg.Command,
		CommandNumber: msg.CommandNumber,
		Stock:         msg.Stock,
		RequestType:   "POST",
	}

	// Checking if user has a valid request array
	if _, ok := UserIdRequestQueueMap[msg.UserId]; !ok {
		UserIdRequestQueueMap[msg.UserId] = &UserTSRequestsDetails{
			UserRequests: &RequestsQueue{
				0, nil, nil,
			}, RunningGoRoutine: false,
		}
	}

	// Adding to the the request queue map
	// glog.Info("Adding to queue")
	UserIdRequestQueueMap[msg.UserId].UserRequests.Enqueue(incomingReq)
	// glog.Info("QUEUES IS:", UserIdRequestQueueMap[msg.UserId].UserRequests.size)
	// UserIdRequestQueueMap[msg.UserId].UserRequests.printQueue()

	// If user doesn't have a routine assigned, start a new one
	if UserIdRequestQueueMap[msg.UserId].RunningGoRoutine == false {
		glog.Info("\n No Routine for this user. Spinning new routine.. ")
		go startProcessingUser(msg.UserId, UserIdRequestQueueMap[msg.UserId].UserRequests)
	}

	// glog.Info("$$$$$$$$ AFTER PROCESSING< QUE UE IS: ")
	// UserIdRequestQueueMap[msg.UserId].UserRequests.printQueue()

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	// w.Write(msgJson)
	// return msg
}

// Should get the request
// Call the appropriate transaction server endpoint
// Passing in all the info in the body

func homeFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to Request Processing Server!")
}

func main() {
	glog.Info("SAVING XML LOG FILE")
	router := mux.NewRouter()
	flag.Usage = usage
	flag.Parse()

	// router.Use(loggingMiddleware)

	// go startProcessing(UserIdRequestQueueMap)

	router.HandleFunc("/", homeFunc).Methods("GET")
	// ##TODO: Implement auth
	router.HandleFunc("/api/authenticate", parseRequest).Methods("POST")
	router.HandleFunc("/api/add", parseRequest).Methods("POST")
	router.HandleFunc("/api/sell", parseRequest).Methods("POST")
	router.HandleFunc("/api/buy", parseRequest).Methods("POST")
	router.HandleFunc("/api/commit_sell", parseRequest).Methods("POST")
	router.HandleFunc("/api/commit_buy", parseRequest).Methods("POST")
	router.HandleFunc("/api/cancel_buy", parseRequest).Methods("POST")
	router.HandleFunc("/api/cancel_sell", parseRequest).Methods("POST")
	router.HandleFunc("/api/set_buy_amount", parseRequest).Methods("POST")
	router.HandleFunc("/api/set_sell_amount", parseRequest).Methods("POST")
	router.HandleFunc("/api/cancel_set_buy", parseRequest).Methods("POST")
	router.HandleFunc("/api/cancel_set_sell", parseRequest).Methods("POST")
	router.HandleFunc("/api/set_buy_trigger", parseRequest).Methods("POST")
	router.HandleFunc("/api/set_sell_trigger", parseRequest).Methods("POST")
	router.HandleFunc("/api/display_summary", parseRequest).Methods("POST")
	router.HandleFunc("/api/dumplog", parseRequest).Methods("POST")
	router.HandleFunc("/api/quote", parseRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":9091", router))

}
