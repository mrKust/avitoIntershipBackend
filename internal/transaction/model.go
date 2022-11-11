package transaction

import "github.com/jackc/pgx/pgtype"

type Transaction struct {
	ID          int         `json:"id"`
	fromId      string      `json:"from_id"`
	ToId        string      `json:"to_id"`
	ForService  string      `json:"for_service"`
	OrderId     string      `json:"order_id"`
	MoneyAmount string      `json:"money_amount"`
	Status      string      `json:"status"`
	Date        pgtype.Date `json:"date"`
}
