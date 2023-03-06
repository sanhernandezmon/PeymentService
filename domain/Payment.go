package domain

type CreatePaymentRequest struct {
	OrderID    string `json:"order_id"`
	TotalPrice int64  `json:"total_price"`
}

type Payment struct {
	PaymentId  string `json:"payment_id"`
	OrderID    string `json:"order_id"`
	TotalPrice int64  `json:"total_price"`
}
