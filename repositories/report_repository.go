package repositories

import (
	"categories-api/models"
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReport(startDate, endDate string) (*models.ReportResponse, error) {

	var report models.ReportResponse

	// total revenue & transaksi
	query1 := `
	SELECT 
		COALESCE(SUM(total_amount),0),
		COUNT(*)
	FROM transactions
	WHERE DATE(created_at) BETWEEN $1 AND $2
	`

	err := r.db.QueryRow(query1, startDate, endDate).
		Scan(&report.TotalRevenue, &report.TotalTransaksi)

	if err != nil {
		return nil, err
	}

	// produk terlaris
	query2 := `
	SELECT p.name, SUM(td.quantity) qty
	FROM transaction_details td
	JOIN products p ON p.id = td.product_id
	JOIN transactions t ON t.id = td.transaction_id
	WHERE DATE(t.created_at) BETWEEN $1 AND $2
	GROUP BY p.name
	ORDER BY qty DESC
	LIMIT 1
	`

	err = r.db.QueryRow(query2, startDate, endDate).
		Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)

	if err != nil {
		return nil, err
	}

	return &report, nil
}
