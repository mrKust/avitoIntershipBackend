package service

import "context"

type Repository interface {
	Create(ctx context.Context, service *Service) error
	FindAll(ctx context.Context) ([]Service, error)
	FindOne(ctx context.Context, id string) (Service, error)
	Update(ctx context.Context, service Service) error
	Delete(ctx context.Context, id string) error
}
