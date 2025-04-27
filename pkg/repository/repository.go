package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/TeamDev/pkg"
)

type IUserRepo interface {
	RegisterUser(user string) (string, error)
}

type Repository struct {
	IUserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo:          NewUserPostgres(db),
	}
}

