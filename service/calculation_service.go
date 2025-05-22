package service

import (
	"errors"
	"net/http"
	"sistem-pembiayaan/dto/calculation"
	"sistem-pembiayaan/repository"
)

type CalculationService interface {
	CalculateMargin(amount int) (*calculation.MarginCalculationResponse, int, error)
}

type calculationService struct {
	tenorRepo repository.TenorRepository
}

func NewCalculationService(tenorRepo repository.TenorRepository) CalculationService {
	return &calculationService{tenorRepo}
}

func (s *calculationService) CalculateMargin(amount int) (*calculation.MarginCalculationResponse, int, error) {
	// Set 20% persentase
	const annualRate = 0.2

	// Get All data tenor
	tenors, err := s.tenorRepo.GetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("gagal mengambil data tenor")
	}

	// Check jika tenor tidak ada
	if len(tenors) == 0 {
		return nil, http.StatusBadRequest, errors.New("data tenor tidak tersedia")
	}

	// Kalkulasi semua tenor dengan rumus perhitungan
	var results []calculation.MarginCalculationItem
	for _, tenor := range tenors {
		t := float64(tenor.TenorValue)
		margin := (float64(amount) * annualRate * t) / 12
		total := float64(amount) + margin
		monthly := total / t

		results = append(results, calculation.MarginCalculationItem{
			Tenor:              tenor.TenorValue,
			MonthlyInstallment: monthly,
			TotalMargin:        margin,
			TotalPayment:       total,
		})
	}

	response := &calculation.MarginCalculationResponse{
		Calculations: results,
	}
	return response, http.StatusOK, nil
}
