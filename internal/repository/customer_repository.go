package repository

import (
	"be-border-service/internal/model"
	"context"
	"database/sql"
)

type customerRepostiory struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepostiory{db: db}
}

func (c *customerRepostiory) Create(ctx context.Context, customer *model.Customer) error {
	return nil
}
func (c *customerRepostiory) FindOne(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	return nil, nil
}
func (c *customerRepostiory) Update(ctx context.Context, customer *model.Customer) error {
	return nil
}
func (c *customerRepostiory) Delete(ctx context.Context, customerID string) error {
	return nil
}
func (c *customerRepostiory) FindAll(ctx context.Context) ([]model.Customer, error) {
	return nil, nil
}
