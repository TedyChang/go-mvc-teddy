package service

import (
	"codetest/security"
	"codetest/user/dto"
	"codetest/user/entity"
	"codetest/user/repository"
	"errors"
	"log"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}

func (r UserService) Login(dto dto.UserLoginDto) (string, error) {
	var user = r.Repository.Login(dto.Email, dto.Password)
	if user.ID == 0 {
		return "", errors.New("repository.login fail")
	}
	encodeUser, err := security.EncodeUser(user.ID)
	if err != nil {
		log.Printf("service - login err : %v", err)
		return "", err
	}
	return *encodeUser, nil
}

func (r UserService) MyInfo(token string) (*security.Claim, error) {
	return security.Decode(token)
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
