package service

import (
	// "github.com/Mamvriyskiy/database_course/main/pkg"
	// "github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type IUser interface {
	Register(user pkg.UserHandler) (string, error)
}

type Services struct {
	IUser
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IUser:          NewUserService(repo.IUserRepo),
	}
}