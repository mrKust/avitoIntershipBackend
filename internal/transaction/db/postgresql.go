package db

import (
	"avitoIntershipBackend/internal/transaction"
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

func (r *repository) Create(ctx context.Context, transaction *transaction.Transaction) error {
	q := `INSERT INTO transaction (from_id, to_id, for_service, order_id, money_amount, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s"), q)
	row := r.client.QueryRow(ctx, q, transaction.FromId, transaction.ToId, transaction.ForService, transaction.OrderId, transaction.MoneyAmount, transaction.Status)
	if err := row.Scan(&transaction.ID); err != nil {
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

func (r *repository) FindAll(ctx context.Context) (t []transaction.Transaction, err error) {
	q := `SELECT id, from_id, to_id, for_service, order_id, money_amount, status, date FROM transaction;`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	transactions := make([]transaction.Transaction, 0)

	for rows.Next() {
		var t transaction.Transaction

		err = rows.Scan(&t.ID, &t.FromId, &t.ToId, &t.ForService, &t.OrderId, &t.MoneyAmount, &t.Status, &t.Date)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (transaction.Transaction, error) {
	q := `SELECT id, from_id, to_id, for_service, order_id, money_amount, status, date FROM transaction WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var t transaction.Transaction
	row := r.client.QueryRow(ctx, q, id)
	err := row.Scan(&t.ID, &t.FromId, &t.ToId, &t.ForService, &t.OrderId, &t.MoneyAmount, &t.Status, &t.Date)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return t, nil
}

func (r *repository) Update(ctx context.Context, transaction transaction.Transaction) error {
	q := `UPDATE transaction SET from_id = $2, to_id = $3, for_service = $4, order_id = $5, money_amount = $6, status = $7, date = $8 WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Query(ctx, q, transaction.FromId, transaction.ToId, transaction.ForService, transaction.OrderId, transaction.MoneyAmount, transaction.Status, transaction.Date)
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
	q := `DELETE FROM transaction WHERE id = $1;`
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

func NewRepository(client postgresql.Client, logger *logging.Logger) transaction.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
