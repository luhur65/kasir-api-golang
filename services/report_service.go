package services

import (
	"api-kasir/models"
	"api-kasir/repositories"
	// "fmt"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.DailyReport, error) {
	now := time.Now()
	tgldari := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tglsampai := tgldari.Add(24 * time.Hour)
	// fmt.Println(tgldari, tglsampai)

	totalRevenue, err := s.repo.GetTotalRevenue(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	totalTransaksi, err := s.repo.GetTotalTransaksi(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	produkTerlaris, err := s.repo.GetProdukTerlaris(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	return &models.DailyReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}, nil
}

func (s *ReportService) GetReportByRange(tgldari, tglsampai time.Time) (*models.DailyReport, error) {
	totalRevenue, err := s.repo.GetTotalRevenue(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	totalTransaksi, err := s.repo.GetTotalTransaksi(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	produkTerlaris, err := s.repo.GetProdukTerlaris(tgldari, tglsampai)
	if err != nil {
		return nil, err
	}

	return &models.DailyReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}, nil
}
