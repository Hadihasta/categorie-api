package services

import (
	"categories-api/models"
	"categories-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(startDate, endDate string) (*models.ReportResponse, error) {
	return s.repo.GetReport(startDate, endDate)
}
