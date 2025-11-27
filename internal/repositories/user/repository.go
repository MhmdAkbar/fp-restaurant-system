package repoitories

import usermodels "aplikasi_restoran/internal/models/user"

type UserRepository interface { // definisi interface repository
	Create(user *usermodels.User) error                 // buat user
	FindByEmail(email string) (*usermodels.User, error) // cari user by email
	FindByID(id uint) (*usermodels.User, error)         // cari user by id
	Update(user *usermodels.User) error
	Delete(user *usermodels.User) error
}