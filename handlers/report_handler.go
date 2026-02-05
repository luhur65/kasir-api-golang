package handlers

import (
	"encoding/json"
	"net/http"
	"api-kasir/services"
	"api-kasir/models"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetDailyReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(report); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	dari := r.URL.Query().Get("start_date")
	sampai := r.URL.Query().Get("end_date")

	var report *models.DailyReport
	var err error

	if dari == "" || sampai == "" {
		report, err = h.service.GetDailyReport()
	} else {
		formatTgl := "2006-01-02"

		tgldari, err1 := time.Parse(formatTgl, dari)
		tglsampai, err2 := time.Parse(formatTgl, sampai)

		if err1 != nil || err2 != nil {
			http.Error(w, "format tanggal harus" + formatTgl, http.StatusBadRequest)
			return
		}

		tglsampai = tglsampai.Add(24 * time.Hour)

		report, err = h.service.GetReportByRange(tgldari, tglsampai)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
