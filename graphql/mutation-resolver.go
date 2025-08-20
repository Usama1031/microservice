package main

import (
	"context"
	"errors"
	"log"
	"time"
)

var (
	ErrInvalidParameter = errors.New("Invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostAccount(ctx, in.Name)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Account{
		ID:   a.ID,
		Name: a.Name,
	}, nil
}

func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.catalogClient.PostProduct(ctx, in.Name, in.Description, in.Price)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Product{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Price:       a.Price,
	}, nil
}

func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
}
