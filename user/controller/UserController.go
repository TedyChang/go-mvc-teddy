package controller

import (
	"codetest/user/dto"
	"codetest/user/service"
	"fmt"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(impl service.UserService) UserController {
	return UserController{impl}
}

func (c UserController) SaveUser(name string) {
	dto1 := dto.SaveUserDto{Name: name}
	c.UserService.Save(dto1)
}

func (c UserController) GetAll() {
	arr := c.UserService.GetAll()
	for _, value := range arr {
		fmt.Println(value)
	}
}

func (c UserController) GetById(id int64) {
	fmt.Println(c.UserService.GetById(id))
}
