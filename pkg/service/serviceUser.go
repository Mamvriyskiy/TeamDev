package service

import (
	"fmt"
	"github.com/Mamvriyskiy/TeamDev/pkg/repository"
)

const (
	salt       = "hfdjmaxckdk20"
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

type UserService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{repo: repo}
}


func (r *UserService) RegisterUser(user string) (id string, err error) {
	fmt.Println("2")
	r.repo.RegisterUser("adfds")
	return id, err
}
