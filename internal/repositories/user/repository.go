package repoitories

import "aplikasi_restoran/internal/models"

type UserRepository interface { // definisi interface repository
	Create(user *models.User) error                 // buat user
	FindByEmail(email string) (*models.User, error) // cari user by email
	FindByID(id uint) (*models.User, error)         // cari user by id
	Update(user *models.User) error
	Delete(user *models.User) error
}
