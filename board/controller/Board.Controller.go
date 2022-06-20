package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
		boardDto := &dto.SaveBoardDto{}

		err := c.ShouldBindJSON(boardDto)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"message": "error : not SaveBoard"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": r.BoardService.Save(*boardDto),
		})
	}
}

func (r Controller) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.BoardService.GetAll()

		c.JSON(http.StatusOK, gin.H{
			"data": arr,
		})
	}
}

func (r Controller) GetById() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"message": "error : id is not number"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": r.BoardService.GetById(id),
		})
	}
}
