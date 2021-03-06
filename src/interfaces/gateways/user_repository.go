package gateways

import (
	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/infrastructure"
)

type UserRepository struct {
	infrastructure.SqlHandler
}

func NewUserRepository(handler *infrastructure.SqlHandler) *UserRepository {
	return &UserRepository{SqlHandler: *handler}
}

func (repo *UserRepository) Persist(u *entities.User) (*entities.User, error) {
	user := entities.User{}
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindById(id int) (*entities.User, error) {
	user := entities.User{}
	err := repo.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindAll() (*entities.Users, error) {
	users := entities.Users{}
	err := repo.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *UserRepository) Update(id int, u *entities.User) (*entities.User, error) {
	user := entities.User{}
	user.ID = id
	user.FirstName = u.FirstName
	user.LastName = u.LastName

	data := entities.User{}
	err := repo.DB.Where("id = ?", id).First(&data).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Delete(id int) error {
	err := repo.DB.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
