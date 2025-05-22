package calculation

type MarginCalculationRequest struct {
	Amount int `json:"amount" validate:"required,gt=0"`
}
