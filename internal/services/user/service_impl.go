package userservice

import (
	helpers "aplikasi_restoran/internal/helper"
	usermodels "aplikasi_restoran/internal/models/user"
	repositories "aplikasi_restoran/internal/repositories/user"

	"errors"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(name, email, pass string, role usermodels.UserRole) error {

	hashed, err := helpers.HashPassword(pass)
	if err != nil {
		return err
	}

	user := usermodels.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     role,
	}

	return s.repo.Create(&user)
}

func (s *userService) Login(email, pass string) (*usermodels.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	if !helpers.CheckPasswordHash(user.Password, pass) {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

func (s *userService) GetProfile(id uint) (*usermodels.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateProfile(id uint, name, email string) (*usermodels.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteProfile (id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}

