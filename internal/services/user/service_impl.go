package userservice

import (
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	repositories "aplikasi_restoran/internal/repositories/user"
	"errors"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(name, email, pass string, role models.UserRole) (*models.User ,error) {

	hashed, err := helpers.HashPassword(pass)
	if err != nil {
		return nil,err
	}
user := models.User{
        Name:     name,
        Email:    email,
        Password: hashed,
        Role:     role,
    }

    if err := s.repo.Create(&user); err != nil {
        return nil, err
    }

    return &user, nil}

func (s *userService) Login(email, pass string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	if !helpers.CheckPasswordHash(user.Password, pass) {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

func (s *userService) GetProfile(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateProfile(id uint, name, email string) (*models.User, error) {
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