package user

import "database/sql"

type UserRepositoryInterface interface {
	GetUser(id int) (*User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUser(id int) (*User, error) {
	return &User{
		Id:   id,
		Name: "User Name",
	}, nil
}
