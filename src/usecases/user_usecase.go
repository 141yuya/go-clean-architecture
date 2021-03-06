package usecases

import (
	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/domain/repositories"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return UserUsecase{userRepository: userRepository}
}

func (userUsecase *UserUsecase) Add(u *entities.User) (*entities.User, error) {
	user, err := userUsecase.userRepository.Persist(u)
	return user, err
}

func (userUsecase *UserUsecase) GetUser(id int) (*entities.User, error) {
	user, err := userUsecase.userRepository.FindById(id)
	return user, err
}

func (userUsecase *UserUsecase) GetUsers() (*entities.Users, error) {
	users, err := userUsecase.userRepository.FindAll()
	return users, err
}

func (userUsecase *UserUsecase) Update(id int, u *entities.User) (*entities.User, error) {
	user, err := userUsecase.userRepository.Update(id, u)
	return user, err
}

func (userUsecase *UserUsecase) Delete(id int) error {
	return userUsecase.userRepository.Delete(id)
}
