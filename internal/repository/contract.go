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
	Create(ctx context.Context, customer model.Users) error
}

// Read (Find All)
type FindAllCustomer interface {
	FindAll(ctx context.Context) ([]model.Users, error)
}

// Read (Find One)
type FindOneCustomer interface {
	FindOne(ctx context.Context, customer model.Users) (*model.Users, error)
}

// Update
type UpdateCustomer interface {
	Update(ctx context.Context, customer model.Users) error
}

// Delete
type DeleteCustomer interface {
	Delete(ctx context.Context, customerID string) error
}
