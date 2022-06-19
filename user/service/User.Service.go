package service

import (
	"codetest/user/dto"
	"codetest/user/entity"
	"codetest/user/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}

func (r UserService) GetById(id int64) entity.User {
	return r.Repository.FindById(id)
}

func (r UserService) Save(dto dto.SaveUserDto) int64 {
	id := r.Repository.Save(entity.User{
		Name: dto.Name,
	})

	return id
}

func (r UserService) GetAll() []entity.User {
	return r.Repository.FindAll()
}
