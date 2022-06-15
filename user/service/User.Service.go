package service

import (
	"codetest/user/entity"
	"codetest/user/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func (r UserServiceImpl) GetById(id int64) {
	r.Repository.FindById(id)
}

func (r UserServiceImpl) Save(id int64, name string) {
	r.Repository.Save(entity.User{
		Id:   id,
		Name: name,
	})
}

func (r UserServiceImpl) GetAll() {
	r.Repository.FindAll()
}
