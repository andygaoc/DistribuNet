package api

import "time"

type Order struct {
	OrderID    string    `json:"order_id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	Timestamp  time.Time `json:"timestamp"`
}
type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Timestamp time.Time `json:"timestamp"`
}
type Bill struct {
	BillID        string    `json:"bill_id"`
	CustomerID    string    `json:"customer_id"`
	Amount        float64   `json:"amount"`
	PaymentStatus string    `json:"payment_status"`
	Timestamp     time.Time `json:"timestamp"`
}

/*
{
"order_id": "1234567890",
"customer_id": "9876543210",
"product_id": "ABC123",
"quantity": 2,
"price": 19.99,
"timestamp": "2022-01-01T10:00:00Z"
}
{
"order_id": "1234567890",
"customer_id": "9876543210",
"product_id": "ABC123",
"quantity": 2,
"price": 19.99,
"timestamp": "2022-01-01T10:00:00Z"
}
{
"bill_id": "9876543210",
"customer_id": "1234567890",
"amount": 99.99,
"payment_status": "paid",
"timestamp": "2022-01-01T10:00:00Z"
}
*/
