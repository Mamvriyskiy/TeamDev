package repository

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	CreateUser(user pkg.UserService) (string, error)
}

type Repository struct {
	IUserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo:          NewUserPostgres(db),
	}
}
