package service

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/gabeefranco/gonotes-api/internal/domain"
	"github.com/gabeefranco/gonotes-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordTooShort = errors.New("password too short (min: 6 characters)")
var ErrInternal = errors.New("internal server error")
var ErrUserAlreadyExists = errors.New("user already exists")

type UsersService struct {
	Repository repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) *UsersService {
	return &UsersService{
		Repository: repo,
	}
}

func (s UsersService) CreateUser(email string, password string) (*domain.User, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmail
	}

	if len(password) < 6 {
		return nil, ErrPasswordTooShort
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	exists, err := s.Repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, ErrInternal
	}

	if exists != nil {
		return nil, ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}

	u := domain.User{
		Email:    email,
		Password: string(passwordHash),
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = s.Repository.Create(ctx, &u)

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}

	return &u, nil
}
