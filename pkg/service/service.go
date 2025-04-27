package service

import (
	// "github.com/Mamvriyskiy/database_course/main/pkg" 
	"github.com/Mamvriyskiy/TeamDev/pkg/repository"
)

type IUser interface {
	RegisterUser(user string) (string, error)
}

type Services struct {
	IUser
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IUser:          NewUserService(repo.IUserRepo),
	}
}