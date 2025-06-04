package repository

import (
	"be-border-service/internal/model"
	"be-border-service/pkg/databasex"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type customerRepostiory struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepostiory{db: db}
}

func (c *customerRepostiory) Create(ctx context.Context, customer model.Users) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := c.db.ExecContext(ctx, query, customer.Name, customer.Email)
	if err != nil {
		return err
	}
	return nil
}
func (c *customerRepostiory) FindOne(ctx context.Context, customer model.Users) (*model.Users, error) {
	var cust model.Users
	q := "SELECT id, name, email FROM users"

	whereClause, args := databasex.BuildWhereClause(customer)

	query := q + whereClause

	err := c.db.QueryRowContext(ctx, query, args...).Scan(&cust.ID, &cust.Name, &cust.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &cust, nil
}

func (c *customerRepostiory) Update(ctx context.Context, customer model.Users) error {
	return nil
}
func (c *customerRepostiory) Delete(ctx context.Context, customerID string) error {
	return nil
}
func (c *customerRepostiory) FindAll(ctx context.Context) ([]model.Users, error) {
	var (
		users []model.Users
	)
	err := c.db.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}
