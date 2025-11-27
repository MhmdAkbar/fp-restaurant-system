package repoitories

import (
	usermodels "aplikasi_restoran/internal/models/user"

	"gorm.io/gorm"
)

type userRepo struct { // struct implementasi interface
	db *gorm.DB // koneksi gorm
}

func NewUserRepository(db *gorm.DB) UserRepository { // constructor
	return &userRepo{db: db} // buat instance repo
}

func (r *userRepo) Create(user *usermodels.User) error { // implement create
	return r.db.Create(user).Error // insert ke database
}

func (r *userRepo) FindByEmail(email string) (*usermodels.User, error) { // implement find by email
	var user usermodels.User                                     // variable user
	err := r.db.Where("email = ?", email).First(&user).Error // query
	return &user, err                                        // kembalikan hasil
}

func (r *userRepo) FindByID(id uint) (*usermodels.User, error) { // implement find by id
	var user usermodels.User               // variable user
	err := r.db.First(&user, id).Error // query
	return &user, err                  // kembalikan hasil
}

func (r *userRepo) Update(user *usermodels.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) Delete(user *usermodels.User) error {
	return r.db.Delete(user).Error
}

