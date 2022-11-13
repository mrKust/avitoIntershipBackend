package transaction

import "context"

type Repository interface {
	Create(ctx context.Context, transaction *Transaction) error
	FindAll(ctx context.Context) ([]Transaction, error)
	FindAllForPeriod(ctx context.Context, month string, year string) ([]Transaction, error)
	FindOne(ctx context.Context, id string) (Transaction, error)
	FindPageForUser(ctx context.Context, id, pageNum, sortSum, sortDate string) ([]Transaction, error)
	Update(ctx context.Context, transaction Transaction) error
	Delete(ctx context.Context, id string) error
}
