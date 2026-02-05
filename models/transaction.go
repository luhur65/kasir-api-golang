package models

type Transaction struct {
	ID         	int     `json:"id"`
	TotalAmount float64 `json:"total_amount"`
	Details 		[]TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID         		int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ProductID   	int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Quantity      int     `json:"quantity"`
	SubTotal    	float64 `json:"subtotal"`
}

type CheckoutRequest struct {
	Items []CheckoutItems `json:"items"`
}

type CheckoutItems struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}