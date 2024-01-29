//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"database/sql"

	"github.com/RomeroGabriel/dependency-inversion/product"
	"github.com/RomeroGabriel/dependency-inversion/user"
	"github.com/google/wire"
)

// Provider: a function that can produce a value.
// Sets are useful if several providers will frequently be used together.
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
	user.NewUserRepository,
	wire.Bind(new(user.UserRepositoryInterface), new(*user.UserRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}

func NewUserUseCase(db *sql.DB) *user.UserUseCase {
	wire.Build(
		setRepositoryDependency,
		user.NewUserUseCase,
	)
	return &user.UserUseCase{}
}
