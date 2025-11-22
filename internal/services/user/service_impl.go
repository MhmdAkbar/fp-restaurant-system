package userservice

import (
	"aplikasi_restoran/internal/models"
	repoitories "aplikasi_restoran/internal/repositories/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct { // implementasi interface
	repo repoitories.UserRepository
}

func NewUserService(repo repoitories.UserRepository) UserService { // return interface
	return &userService{repo}
}

func (s *userService) Register(name, email, pass string, role models.UserRole) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		Role:     role,
	}

	return s.repo.Create(&user)
}

func (s *userService) Login(email, pass string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) != nil {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

func (s *userService) GetProfile(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}
