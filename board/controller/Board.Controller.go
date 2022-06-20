package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"codetest/util/gin_util"
	"github.com/gin-gonic/gin"
	"strconv"
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
			gin_util.BadJson(c, gin.H{"message": "error : not SaveBoard"})
			return
		}

		gin_util.OkJson(c, gin.H{"id": r.BoardService.Save(boardDto)})
	}
}

func (r Controller) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.BoardService.GetAll()

		gin_util.OkJson(c, gin.H{"data": arr})
	}
}

func (r Controller) GetById() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err1 := strconv.ParseInt(c.Param("id"), 10, 64)
		if err1 != nil {
			gin_util.BadJson(c, gin.H{"message": "error : id is not number"})
			return
		}

		gin_util.OkJson(c, gin.H{"data": r.BoardService.GetById(id)})
	}
}
