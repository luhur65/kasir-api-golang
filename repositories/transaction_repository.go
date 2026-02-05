package repositories

import (
	"database/sql"
	"errors"
	"api-kasir/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItems) (*models.Transaction, error) {
	var (
		res *models.Transaction
	)

	tx, err := repo.db.Begin()
	if  err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// inisialisasi total amount
	var totalAmount int64 = 0

	details := make([]models.TransactionDetail, 0)

	// hitung total amount dan buat detail transaksi
	for _, item := range items {
		var ProductName string
		var productID, price, stock int64
		err := tx.QueryRow(`SELECT id, name, price, stock FROM products WHERE id = $1`, item.ProductID).Scan(&productID, &ProductName, &price, &stock)

		if err == sql.ErrNoRows {
			return nil, errors.New("product not " + string(rune(item.ProductID)) + " found")
		}

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		subtotal := price * int64(item.Quantity)
		totalAmount += subtotal

		_, err = tx.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2`, item.Quantity, item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   int(productID),
			ProductName: ProductName,
			Quantity:    item.Quantity,
			SubTotal:    float64(price * int64(item.Quantity)),
		})
	}

	var transactionID int64
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i, detail := range details {
		details[i].TransactionID = int(transactionID)
		_, err = tx.Exec(
			`INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal)
			VALUES ($1, $2, $3, $4)`,
			transactionID,
			detail.ProductID,
			detail.Quantity,
			detail.SubTotal,
		)
		if err != nil {
			return nil, err
		}

	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	res = &models.Transaction{
		ID:          int(transactionID),
		TotalAmount: float64(totalAmount),
		Details:     details,
	}

	return res, nil;
	
}