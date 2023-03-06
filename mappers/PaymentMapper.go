package mappers

import (
	"PeymentService/domain"
	"github.com/google/uuid"
)

func MapPaymentRequestToPayment(request domain.CreatePaymentRequest) domain.Payment {
	newUUID := uuid.New().String()
	return domain.Payment{newUUID, request.OrderID, request.TotalPrice}
}
