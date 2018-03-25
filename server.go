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

// var UserIdRequestQueueMap = make(map[string][]TSRequest)
var UserIdRequestQueueMap = make(map[string]UserTSRequestsDetails)

// var arr = *NewQueue(1)

// var UserIdRequestQueueMap = &UserTSRequestsDetails{UserRequests: make([]TSRequest, 0)}
// var UserIdRequestQueueMap = make(map[string]UserTSRequestsDetails) .UserRequests{UserRequests: make([]TSRequest, 0)})

var transactionHistory []string

// func authenticateUser(userId string) {
// 	f = getFilePointer()
// account := initializeAccount(userId)
// UserMap[userId] = &account
// glog.Info("##### Account Balance: ", account.Balance, " Available: ", account.Available)
// glog.Info("Account Stocks: ", account.StockPortfolio["S"])
// 	glog.Info("INFO: Retrieving user from the db..")
// }

// type TSRequest struct {
// 	UserId        string
// 	PriceDollars  float64
// 	PriceCents    float64
// 	Command       string
// 	CommandNumber int
// 	Stock         string
// 	RequestType   string
// }

type UserTSRequestsDetails struct {
	UserRequests *RequestsQueue
	// UserRequests     RequestsQueue{nodes: make([]TSRequest, 0)}
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

// func startProcessing(UserIdRequestQueueMap map[string]UserTSRequestsDetails, userId string) {
// 	if UserIdRequestQueueMap[userId].RunningGoRoutine == false {
// 		go constructAndRequestTS(userId, UserIdRequestQueueMap[userId].UserRequests)

// 		for user, requestQueue := range UserIdRequestQueueMap {
// 			glog.Info("LOOKIN AT USER:", user)

// 		}

// 	}
// }

// func constructAndRequestTS(userId string, userRequestsQueue *[]TSRequest) {
// 	glog.Info("\n \n \n \n  inside routinee: user   ", userId)

// 	for index, request := range *userRequestsQueue {
// 		// element is the element from someSlice for where we are
// 		// processRequest()
// 		glog.Info("removing: ", index)
// 		*userRequestsQueue = *userRequestsQueue[1:]
// 	}
// }

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

//Parse request and return Response Object
// Get the user request -- done
// Add to user map -- done
// If there is a go routine thats servicing a user, do nothing
// If there is no go routine, start one up
// In the go routine, create the hhtp request and fire
func parseRequest(w http.ResponseWriter, r *http.Request) {
	msg := Response{}
	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
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
		Stock:         "string",
		RequestType:   "POST",
	}

	// Adding requests into queue
	glog.Info("\n \n\n\n @@@@@@ THE QUEUE IS:", msg.UserId, UserIdRequestQueueMap[msg.UserId].UserRequests)
	// if UserIdRequestQueueMap[msg.UserId] == nil { //.count < 1 {
	if _, ok := UserIdRequestQueueMap[msg.UserId]; !ok {
		glog.Info("\n %%%% creating a new array for user:", msg.UserId)
		// UserIdRequestQueueMap[msg.UserId] = UserTSRequestsDetails {
		// 	UserRequests: make([]TSRequest, 0)
		// }

		// UserIdRequestQueueMap[msg.UserId] = UserTSRequestsDetails{UserRequests: &arr, RunningGoRoutine: false}
		UserIdRequestQueueMap[msg.UserId] = UserTSRequestsDetails{
			UserRequests: &RequestsQueue{
				[]TSRequest{
					{},
				}, 1, 0, 0, 0,
			}, RunningGoRoutine: false,
		}
	}

	UserIdRequestQueueMap[msg.UserId].UserRequests.Push(incomingReq)
	// UserIdRequestQueueMap[msg.UserId].UserRequests.Push(TSRequest{
	// 	UserId:        msg.UserId,
	// 	PriceDollars:  msg.PriceDollars,
	// 	PriceCents:    msg.PriceCents,
	// 	Command:       msg.Command,
	// 	CommandNumber: msg.CommandNumber,
	// 	Stock:         "string",
	// 	RequestType:   "POST",
	// })

	// glog.Info("After pushing: >>>>> ", arr)

	glog.Info("\nNOW THE USER QUEUE IS ", msg.UserId, UserIdRequestQueueMap[msg.UserId].UserRequests)

	//TODO: rewrite this!!
	// switch msg.Command {
	// case "authenticate":
	// 	authenticateUser(msg.UserId)
	// case "add":
	// 	addFunds(msg.UserId, msg.PriceDollars, msg.PriceCents, msg.CommandNumber)
	// 	// add(account, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("SUCCESS: Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("SUCCESS: Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: ADD Successful \n")
	// case "buy":
	// 	glog.Info("\n\n############################### INFO: Executing BUY FOR... ", msg.PriceDollars, msg.CommandNumber)
	// 	// buy(account, msg.Stock, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: BUY Successful")
	// case "sell":
	// 	glog.Info("\n\n############################### INFO: Executing SELL OF ", msg.Stock, msg.CommandNumber)
	// 	// sell(account, msg.Stock,msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: BUY Successful")
	// case "commit_sell":
	// 	glog.Info("\n\n############################### INFO: Executing COMMIT SELL ", msg.CommandNumber)
	// 	// commitSell(account, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: COMMIT SELL Successful")
	// case "commit_buy":
	// 	glog.Info("\n\n############################### INFO: Executing COMMIT BUY ", msg.CommandNumber)
	// 	// commitBuy(account, f, msg.CommandNumber, msg.Command)
	// 	glog.Info("\n############################### SUCCESS: COMMIT BUY Successful")
	// 	// glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// case "cancel_buy":
	// 	glog.Info("\n\n############################### INFO: Executing CANCEL BUY ", msg.CommandNumber)
	// 	// cancelBuy(account, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: CANCEL BUY Successful")
	// case "cancel_sell":
	// 	glog.Info("\n\n############################### INFO: Executing CANCEL SELL ", msg.CommandNumber)
	// 	// cancelSell(account, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: CANCEL SELL Successful")
	// case "set_buy_amount":
	// 	glog.Info("\n\n############################### INFO: Executing SET BUY AMOUNT ", msg.CommandNumber)
	// 	// setBuyAmount(account, msg.Stock, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: SET BUY AMOUNT Successful")
	// case "set_sell_amount":
	// 	glog.Info("\n\n############################### INFO: Executing SET SELL AMOUNT ", msg.CommandNumber)
	// 	// setSellAmount(account, msg.Stock, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: SET SELL AMOUNT Successful")
	// case "cancel_set_buy":
	// 	glog.Info("\n\n############################### INFO: Executing CANCEL SET BUY ", msg.CommandNumber)
	// 	// cancelSetBuy(account, msg.Stock, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: CANCEL SET BUY Successful")
	// case "cancel_set_sell":
	// 	glog.Info("\n\n############################### INFO: Executing CANCEL SET SELL ", msg.CommandNumber)
	// 	// cancelSetSell(account, msg.Stock, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: CANCEL SET SELL Successful")
	// case "set_buy_trigger":
	// 	glog.Info("\n\n############################### INFO: Executing SET BUY TRIGGER ", msg.CommandNumber)
	// 	// setBuyTrigger(account, msg.Stock, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: SET BUY TRIGGER Successful")
	// case "set_sell_trigger":
	// 	glog.Info("\n\n############################### INFO: Executing SET SELL TRIGGER ", msg.CommandNumber)
	// 	// setSellTrigger(account, msg.Stock, msg.PriceDollars, f, msg.CommandNumber, msg.Command)
	// 	// // // glog.Info("Account Balance: ", account.Balance, " Available: ", account.Available)
	// 	// // glog.Info("Account Stocks: ", account.StockPortfolio["S"])
	// 	glog.Info("\n############################### SUCCESS: SET SELL TRIGGER Successful")
	// case "display_summary":
	// 	glog.Info("\n\n############################### INFO: Executing DISPLAY SUMMARY ", msg.CommandNumber)
	// 	// displaySummary(account, transactionHistory)
	// 	glog.Info("\n############################### SUCCESS: DISPLAY SUMMARY Successful")
	// case "dumplog":
	// 	glog.Info("SAVING XML LOG FILE")
	// default:
	// 	panic("Oh noooo we can't process this request :(")
	// }

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

	router.Use(loggingMiddleware)

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
