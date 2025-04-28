package service

import (
	// "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/TeamDev/pkg"
	"github.com/Mamvriyskiy/TeamDev/pkg/repository"
)

type IUser interface {
	RegisterUser(user int) (string, error)
	ProfileUser(user int) (pkg.UserAccount, error)
	AddSocialUser(userID int, url string) error
}

type Services struct {
	IUser
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IUser: NewUserService(repo.IUserRepo),
	}
}
