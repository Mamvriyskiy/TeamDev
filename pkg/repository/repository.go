package repository

import (
	"github.com/Mamvriyskiy/TeamDev/pkg"
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/TeamDev/pkg"
)

type IUserRepo interface {
	RegisterUser(user int) (string, error)
	ProfileUser(user int) (pkg.UserAccount, error)
	AddSocialUser(userID int, url string) error
	CreateTasks(userID int, tasks pkg.NewTasks) (err error)
	CheckStatusTasks(userID int, nameTasks string) (count, all int, err error)
}

type Repository struct {
	IUserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo: NewUserPostgres(db),
	}
}
