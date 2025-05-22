package calculation

type MarginCalculationItem struct {
	Tenor              int     `json:"tenor"`
	MonthlyInstallment float64 `json:"monthly_installment"`
	TotalMargin        float64 `json:"total_margin"`
	TotalPayment       float64 `json:"total_payment"`
}

type MarginCalculationResponse struct {
	Calculations []MarginCalculationItem `json:"calculations"`
}
