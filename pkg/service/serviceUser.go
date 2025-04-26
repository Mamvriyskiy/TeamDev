
import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
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
