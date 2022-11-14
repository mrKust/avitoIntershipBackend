package masterBalance

type MasterBalance struct {
	ID          int    `json:"id"`
	FromId      string `json:"from_id"`
	ServiceId   string `json:"service_id"`
	OrderId     string `json:"order_id"`
	MoneyAmount string `json:"money_amount"`
}

type CreateDTO struct {
	FromId      string `json:"from_id"`
	ServiceId   string `json:"service_id"`
	OrderId     string `json:"order_id"`
	MoneyAmount string `json:"money_amount"`
}
