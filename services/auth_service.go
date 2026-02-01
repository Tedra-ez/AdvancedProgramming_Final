package services

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

type AuthService struct {
	users  *repository.UserRepository
	secret []byte
}

func NewAuthService(users *repository.UserRepository) *AuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev_secret"
	}
	return &AuthService{
		users:  users,
		secret: []byte(secret),
	}
}

func (s *AuthService) Register(ctx context.Context, fullName, email, password string) error {
	email = strings.ToLower(strings.TrimSpace(email))

	_, err := s.users.FindByEmail(ctx, email)
	if err == nil {
		return ErrEmailExists
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &models.User{
		FullName:     fullName,
		Email:        email,
		PasswordHash: string(hash),
		Role:         "customer",
	}

	return s.users.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"sub":  user.ID.Hex(),
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}
