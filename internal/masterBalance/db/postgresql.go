package db

import (
	"avitoIntershipBackend/internal/masterBalance"
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

func (r *repository) Create(ctx context.Context, masterBalance *masterBalance.MasterBalance) error {
	q := `INSERT INTO masterBalance (from_id, service_id, order_id, money_amount) VALUES ($1, $2, $3, $4) RETURNING id`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s"), q)
	row := r.client.QueryRow(ctx, q, masterBalance.FromId, masterBalance.ServiceId, masterBalance.OrderId, masterBalance.MoneyAmount)
	if err := row.Scan(&masterBalance.ID); err != nil {
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

func (r *repository) FindAll(ctx context.Context) (m []masterBalance.MasterBalance, err error) {
	q := `SELECT id, from_id, service_id, order_id, money_amount FROM masterbalance;`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	masterBalances := make([]masterBalance.MasterBalance, 0)

	for rows.Next() {
		var m masterBalance.MasterBalance

		err = rows.Scan(&m.ID, &m.ServiceId, &m.FromId, &m.MoneyAmount, &m.OrderId)
		if err != nil {
			return nil, err
		}

		masterBalances = append(masterBalances, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return masterBalances, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (masterBalance.MasterBalance, error) {
	q := `SELECT id, from_id, service_id, order_id, money_amount FROM masterbalance WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var m masterBalance.MasterBalance
	row := r.client.QueryRow(ctx, q, id)
	err := row.Scan(&m.ID, &m.FromId, &m.OrderId, &m.ServiceId, &m.MoneyAmount)
	if err != nil {
		return masterBalance.MasterBalance{}, err
	}

	return m, nil
}

func (r *repository) Update(ctx context.Context, masterBalance masterBalance.MasterBalance) error {
	q := `UPDATE masterbalance SET from_id = $2, service_id = $3, order_id = $4, money_amount = $5 WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Query(ctx, q, masterBalance.ID, masterBalance.FromId, masterBalance.ServiceId, masterBalance.OrderId, masterBalance.MoneyAmount)
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
	q := `DELETE FROM masterbalance WHERE id = $1;`
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

func NewRepository(client postgresql.Client, logger *logging.Logger) masterBalance.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
