package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Neymand/todo-app"
	"github.com/Neymand/todo-app/pkg/repository"
)

const salt = "fasg4afc3tq534tghg43h5u45etq"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
