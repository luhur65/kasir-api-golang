package repositories

import (
	"database/sql"
	"time"
	// "errors"
	"api-kasir/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetTotalRevenue(tgldari, tglsampai time.Time) (int64, error) {
	query := `
		SELECT COALESCE(SUM(total_amount), 0)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`

	var total int64
	err := repo.db.QueryRow(query, tgldari, tglsampai).Scan(&total)
	return total, err
}

func (repo *ReportRepository) GetTotalTransaksi(tgldari, tglsampai time.Time) (int, error) {
	query := `
		SELECT COUNT(id)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`

	var total int
	err := repo.db.QueryRow(query, tgldari, tglsampai).Scan(&total)
	return total, err
}

func (repo *ReportRepository) GetProdukTerlaris(start, end time.Time) (*models.ProdukTerlaris, error) {
	query := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`

	var produk models.ProdukTerlaris
	err := repo.db.QueryRow(query, start, end).Scan(&produk.Nama, &produk.QtyTerjual)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &produk, nil
}
