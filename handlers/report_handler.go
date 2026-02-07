package handlers

import (
	"categories-api/services"
	"encoding/json"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {

	start := r.URL.Query().Get("start_date")
	end := r.URL.Query().Get("end_date")

	Report, err := h.service.GetReport(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Report)
}

func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")

	report, err := h.service.GetReport(today, today)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(report)
}
