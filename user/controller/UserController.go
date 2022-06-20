package controller

import (
	"codetest/user/dto"
	"codetest/user/service"
	"codetest/util/gu"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(impl service.UserService) UserController {
	return UserController{impl}
}

func (r UserController) SaveUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var userDto dto.SaveUserDto

		err1 := c.ShouldBindJSON(&userDto)
		if err1 != nil {
			gu.BadJson(c, gin.H{"message": "error : not SaveBoard"})
		}

		gu.OkJson(c, gin.H{"id": r.UserService.Save(userDto)})

	}
}

func (r UserController) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.UserService.GetAll()

		gu.OkJson(c, gin.H{"data": arr})
	}
}

func (r UserController) GetById() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err1 := gu.GetID("user", c)
		if err1 != nil {
			gu.BadJson(c, gin.H{"message": "error : not userId"})
			return
		}

		arr := r.UserService.GetById(userId)
		gu.OkJson(c, gin.H{"data": arr})
	}
}
