package db

import (
	"avitoIntershipBackend/internal/service"
	"avitoIntershipBackend/pkg/client/postgresql"
	"avitoIntershipBackend/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, service *service.Service) error {
	q := `INSERT INTO service (name, price) VALUES ($1, $2) RETURNING id`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s"), q)
	row := r.client.QueryRow(ctx, q, service.Name, service.Price)
	if err := row.Scan(&service.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (s []service.Service, err error) {
	q := `SELECT id, name, price FROM service;`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	services := make([]service.Service, 0)

	for rows.Next() {
		var s service.Service

		err = rows.Scan(&s.ID, &s.Name, &s.Price)
		if err != nil {
			return nil, err
		}

		services = append(services, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (service.Service, error) {
	q := `SELECT id, name, price FROM service WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var s service.Service
	row := r.client.QueryRow(ctx, q, id)
	err := row.Scan(&s.ID, &s.Name, &s.Price)
	if err != nil {
		return service.Service{}, err
	}

	return s, nil
}

func (r *repository) Update(ctx context.Context, service service.Service) error {
	q := `UPDATE service SET name = $2, price = $3 WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Query(ctx, q, service.ID, service.Name, service.Price)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM service WHERE id = $1;`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Query(ctx, q, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) service.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
