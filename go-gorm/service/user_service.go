package service

import(
	"go-tutor/go-gorm/model"
	"go-tutor/go-gorm/repository"
)

type UserService interface{
	GetAllUsers() ([]model.User, error)
    GetUserById(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserById(id uint) (*model.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
