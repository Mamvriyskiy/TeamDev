package service

import (
	"fmt"

	"github.com/Mamvriyskiy/TeamDev/pkg"
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

func (r *UserService) RegisterUser(user int) (id string, err error) {
	fmt.Println("2")
	r.repo.RegisterUser(user)
	return id, err
}

func (r *UserService) ProfileUser(userID int) (pkg.UserAccount, error) {
	return r.repo.ProfileUser(userID)
}

func (r *UserService) AddSocialUser(userID int, url string) error {
	return r.repo.AddSocialUser(userID, url)
}
