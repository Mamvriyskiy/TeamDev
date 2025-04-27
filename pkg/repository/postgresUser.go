package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/TeamDev/pkg"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) RegisterUser(user string) (string, error) {
	fmt.Println("uraaaaaaa")

	return "", nil
}
