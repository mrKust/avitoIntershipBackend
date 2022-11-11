package db

import (
	"avitoIntershipBackend/internal/user"
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

func (r *repository) Create(ctx context.Context, user *user.User) error {
	q := `INSERT INTO "user" (balance) VALUES ($1) RETURNING id`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s"), q)
	row := r.client.QueryRow(ctx, q, user.Balance)
	if err := row.Scan(&user.ID); err != nil {
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

func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {
	q := `SELECT id, balance FROM "user";`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var u user.User

		err = rows.Scan(&u.ID, &u.Balance)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (user.User, error) {
	q := `SELECT id, balance FROM "user" WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var u user.User
	row := r.client.QueryRow(ctx, q, id)
	err := row.Scan(&u.ID, &u.Balance)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (r *repository) Update(ctx context.Context, user user.User) error {
	q := `UPDATE "user" SET balance = $2 WHERE id = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Query(ctx, q, user.ID, user.Balance)
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
	q := `DELETE FROM "user" WHERE id = $1;`
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

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
