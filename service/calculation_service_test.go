package service_test

import (
	"errors"
	"sistem-pembiayaan/entity"
	"sistem-pembiayaan/mocks"
	"sistem-pembiayaan/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMargin_Success(t *testing.T) {
	mockRepo := new(mocks.MockTenorRepository)
	service := service.NewCalculationService(mockRepo)

	tenors := []entity.Tenor{
		{TenorValue: 6},
		{TenorValue: 12},
	}

	mockRepo.On("GetAll").Return(tenors, nil)

	amount := 1000000
	resp, status, err := service.CalculateMargin(amount)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Calculations, 2)

	// Contoh cek hasil tenor 6 bulan
	expectedMargin := (float64(amount) * 0.2 * 6) / 12
	expectedTotal := float64(amount) + expectedMargin
	expectedMonthly := expectedTotal / 6

	assert.Equal(t, 6, resp.Calculations[0].Tenor)
	assert.InDelta(t, expectedMonthly, resp.Calculations[0].MonthlyInstallment, 0.001)
	assert.InDelta(t, expectedMargin, resp.Calculations[0].TotalMargin, 0.001)
	assert.InDelta(t, expectedTotal, resp.Calculations[0].TotalPayment, 0.001)
}

func TestCalculateMargin_GetAllError(t *testing.T) {
	mockRepo := new(mocks.MockTenorRepository)
	service := service.NewCalculationService(mockRepo)

	mockRepo.On("GetAll").Return(nil, errors.New("db error"))

	amount := 1000000
	resp, status, err := service.CalculateMargin(amount)

	assert.Nil(t, resp)
	assert.Equal(t, 500, status)
	assert.EqualError(t, err, "gagal mengambil data tenor")
}

func TestCalculateMargin_EmptyTenor(t *testing.T) {
	mockRepo := new(mocks.MockTenorRepository)
	service := service.NewCalculationService(mockRepo)

	mockRepo.On("GetAll").Return([]entity.Tenor{}, nil)

	amount := 1000000
	resp, status, err := service.CalculateMargin(amount)

	assert.Nil(t, resp)
	assert.Equal(t, 400, status)
	assert.EqualError(t, err, "data tenor tidak tersedia")
}
