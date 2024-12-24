package repository

import (
	"go-tutor/go-gorm/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface{
	FindAll() ([]model.User, error)
	FindById(id uint)(*model.User, error)
	Create(user *model.User)error
	Update(user *model.User)error
	Delete(id uint)error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id uint)(*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User)error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *model.User)error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id uint)error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}



