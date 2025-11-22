package repoitories

import (
	"aplikasi_restoran/internal/models"
	"gorm.io/gorm"
)

type userRepo struct { // struct implementasi interface
	db *gorm.DB // koneksi gorm
}

func NewUserRepository(db *gorm.DB) UserRepository { // constructor
	return &userRepo{db: db} // buat instance repo
}

func (r *userRepo) Create(user *models.User) error { // implement create
	return r.db.Create(user).Error // insert ke database
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) { // implement find by email
	var user models.User                                     // variable user
	err := r.db.Where("email = ?", email).First(&user).Error // query
	return &user, err                                        // kembalikan hasil
}

func (r *userRepo) FindByID(id uint) (*models.User, error) { // implement find by id
	var user models.User               // variable user
	err := r.db.First(&user, id).Error // query
	return &user, err                  // kembalikan hasil
}
