package repository

import (
	"be-border-service/internal/model"
	"context"
)

type CustomerRepository interface {
	CreateCustomer
	FindAllCustomer
	FindOneCustomer
	UpdateCustomer
	DeleteCustomer
}

// Create
type CreateCustomer interface {
	Create(ctx context.Context, customer *model.Customer) error
}

// Read (Find All)
type FindAllCustomer interface {
	FindAll(ctx context.Context) ([]model.Customer, error)
}

// Read (Find One)
type FindOneCustomer interface {
	FindOne(ctx context.Context, customer *model.Customer) (*model.Customer, error)
}

// Update
type UpdateCustomer interface {
	Update(ctx context.Context, customer *model.Customer) error
}

// Delete
type DeleteCustomer interface {
	Delete(ctx context.Context, customerID string) error
}
