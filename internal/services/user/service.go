package userservice

import usermodels "aplikasi_restoran/internal/models/user"

type UserService interface { // interface service
	Register(name, email, pass string, role usermodels.UserRole) error
	Login(email, pass string) (*usermodels.User, error)
	GetProfile(id uint) (*usermodels.User, error)
	UpdateProfile(id uint, name, email string) (*usermodels.User, error)
	DeleteProfile(id uint) error
}
