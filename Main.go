package main

import (
	"PeymentService/domain"
	"PeymentService/mappers"
	"PeymentService/repository"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
	consumer "github.com/haijianyang/go-sqs-consumer"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func main() {
	handleRequests()
	repository.RecieveSQSMessages().On(consumer.EventReceiveMessage, consumer.OnReceiveMessage(func(messages []*sqs.Message) {
		for _, s := range messages {
			var paymentRequest domain.CreatePaymentRequest
			json.Unmarshal([]byte(*s.Body), &paymentRequest)
			payment := mappers.MapPaymentRequestToPayment(paymentRequest)
			repository.SendPaymentSQSMessage(payment)
			repository.SendPaymentSQSMessage(payment)
		}
	}))
}
