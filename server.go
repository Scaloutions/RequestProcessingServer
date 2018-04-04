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
	glog.Info("PROCESSING PATH: ", path, " Command number: ", CommandNumber)
	var httpResponse = sendHttpRequest(requestBody, path)
	var httpJsonResponse map[string]interface{}
	json.NewDecoder(httpResponse.Body).Decode(&httpJsonResponse)

	if httpResponse.StatusCode == 200 {
		glog.Info("Request was successful. Response is:", httpJsonResponse)
	} else {
		// TODO: WHAT TO DO HERE?
		glog.Info(">>>>>>>>>>>>> OOPS :(")
	}

	// glog.Info("######### BEFORE REMOVING")
	// RequestsQueue.printQueue()
	RequestsQueue.Dequeue()
	// glog.Info("######### AFTER REMOVING")
	// RequestsQueue.printQueue()

}

func startProcessingUser(userId string, RequestsQueue *RequestsQueue) {
	UserIdRequestQueueMap[userId].RunningGoRoutine = true
	glog.Info("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ Inside a new routine for user: ", userId)
	currentReqNode := RequestsQueue.Head()

	for currentReqNode != nil {
		i := currentReqNode.value
		currentRequest := i.(TSRequest)

		switch currentRequest.Command {
		case ADD:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, ADD, RequestsQueue, currentRequest.CommandNumber)
			break

		case QUOTE:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, QUOTE, RequestsQueue, currentRequest.CommandNumber)
			break

		case BUY:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, BUY, RequestsQueue, currentRequest.CommandNumber)
			break

		case COMMIT_BUY:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, COMMIT_BUY, RequestsQueue, currentRequest.CommandNumber)
			break

		case CANCEL_BUY:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, CANCEL_BUY, RequestsQueue, currentRequest.CommandNumber)
			break

		case SELL:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, SELL, RequestsQueue, currentRequest.CommandNumber)
			break

		case COMMIT_SELL:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, COMMIT_SELL, RequestsQueue, currentRequest.CommandNumber)
			break

		case CANCEL_SELL:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, CANCEL_SELL, RequestsQueue, currentRequest.CommandNumber)
			break

		case SET_BUY_AMOUNT:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, SET_BUY_AMOUNT, RequestsQueue, currentRequest.CommandNumber)
			break

		case CANCEL_SET_BUY:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, CANCEL_SET_BUY, RequestsQueue, currentRequest.CommandNumber)
			break

		case SET_BUY_TRIGGER:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, SET_BUY_TRIGGER, RequestsQueue, currentRequest.CommandNumber)
			break

		case SET_SELL_AMOUNT:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, SET_SELL_AMOUNT, RequestsQueue, currentRequest.CommandNumber)
			break

		case SET_SELL_TRIGGER:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"PriceDollars":  currentRequest.PriceDollars,
				"PriceCents":    currentRequest.PriceCents,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, SET_SELL_TRIGGER, RequestsQueue, currentRequest.CommandNumber)
			break

		case CANCEL_SET_SELL:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"Stock":         currentRequest.Stock,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, CANCEL_SET_SELL, RequestsQueue, currentRequest.CommandNumber)
			break

		case DISPLAY_SUMMARY:
			requestBody := map[string]interface{}{
				"UserId":        userId,
				"CommandNumber": currentRequest.CommandNumber}

			// should call audit server
			handleHttpRequestResponse(requestBody, DISPLAY_SUMMARY, RequestsQueue, currentRequest.CommandNumber)
			break

		case DUMPLOG:
			requestBody := map[string]interface{}{
				"FileName":      userId,
				"CommandNumber": currentRequest.CommandNumber}

			handleHttpRequestResponse(requestBody, DUMPLOG, RequestsQueue, currentRequest.CommandNumber)
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
		glog.Info("\n No Routine for this user: ", msg.UserId, ". Spinning new routine.. \n\n\n\n\n\n")
		go startProcessingUser(msg.UserId, UserIdRequestQueueMap[msg.UserId].UserRequests)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to Request Processing Server!")
}

func main() {
	glog.Info("SAVING XML LOG FILE")
	router := mux.NewRouter()
	flag.Usage = usage
	flag.Parse()

	// router.Use(loggingMiddleware)

	router.HandleFunc("/", homeFunc).Methods("GET")
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
