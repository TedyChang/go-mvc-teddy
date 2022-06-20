package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"codetest/util/gu"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	BoardService boardService.BoardService
}

func NewController(impl boardService.BoardService) Controller {
	return Controller{impl}
}

func (r Controller) SaveBoard() func(c *gin.Context) {
	return func(c *gin.Context) {
		var boardDto dto.SaveBoardDto

		err1 := c.ShouldBindJSON(&boardDto)
		if err1 != nil {
			gu.BadJson(c, gin.H{"message": "error : not SaveBoard"})
			return
		}

		gu.OkJson(c, gin.H{"id": r.BoardService.Save(boardDto)})
	}
}

func (r Controller) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.BoardService.GetAll()

		gu.OkJson(c, gin.H{"data": arr})
	}
}

func (r Controller) GetById() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err1 := gu.GetID("id", c)
		if err1 != nil {
			gu.BadJson(c, gin.H{"message": "error : id is not number"})
			return
		}

		gu.OkJson(c, gin.H{"data": r.BoardService.GetById(id)})
	}
}
