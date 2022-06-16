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
	var arr []entity.User
	arr = r.Repository.FindAll()

	var max = int64(0)
	for _, v := range arr {
		if v.Id > max {
			max = v.Id
		}
	}

	id := r.Repository.Save(entity.User{
		Id:   max + 1,
		Name: dto.Name,
	})

	return id
}

func (r UserService) GetAll() []entity.User {
	return r.Repository.FindAll()
}
