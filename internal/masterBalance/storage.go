package masterBalance

import "context"

type Repository interface {
	Create(ctx context.Context, masterBalance *MasterBalance) error
	FindAll(ctx context.Context) ([]MasterBalance, error)
	FindOne(ctx context.Context, id string) (MasterBalance, error)
	FindOneByParam(ctx context.Context, masterBalance *MasterBalance) error
	Update(ctx context.Context, masterBalance MasterBalance) error
	Delete(ctx context.Context, id string) error
}
