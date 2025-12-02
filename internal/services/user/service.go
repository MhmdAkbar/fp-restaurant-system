package userservice

import "aplikasi_restoran/internal/models"

type UserService interface { // interface service
	Register(name, email, pass string, role models.UserRole) (*models.User, error)
	Login(email, pass string) (*models.User, error)
	GetProfile(id uint) (*models.User, error)
	UpdateProfile(id uint, name, email string) (*models.User, error)
	DeleteProfile(id uint) error
}