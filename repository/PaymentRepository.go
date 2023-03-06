package repository

import (
	"PeymentService/domain"
	"PeymentService/mappers"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	consumer "github.com/haijianyang/go-sqs-consumer"
)

func SavePaymentToDynamoDB(request domain.CreatePaymentRequest) (string, error) {
	var payment = mappers.MapPaymentRequestToPayment(request)
	err := AddElement(payment)
	if err != nil {
		panic(err)
		return "", err
	}
	return payment.PaymentId, err
}

func SendPaymentSQSMessage(payment domain.Payment) {
	message, err := json.Marshal(payment)
	if err != nil {
		panic(err)
	}
	queueURL := "http://localhost:9324/queue/payments"
	sqsURL := "http://localhost:9324"
	sqsClient := newSQS(endpoints.UsEast1RegionID, sqsURL)
	print("sending message to sqs")
	sendMessage(sqsClient, string(message), queueURL)
}

func RecieveSQSMessages() *consumer.Worker {
	queueURL := "http://localhost:9324/queue/orders"
	worker := consumer.New(&consumer.Config{
		Region:   aws.String(endpoints.UsEast1RegionID),
		QueueUrl: aws.String(queueURL),
	}, nil)
	return worker
}
