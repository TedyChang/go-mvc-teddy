package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	BoardService boardService.BoardService
}

func NewController(impl boardService.BoardService) Controller {
	return Controller{impl}
}

func (r Controller) SaveBoard(title string, contents string) {
	dto1 := dto.SaveBoardDto{
		Title:    title,
		Contents: contents,
	}

	r.BoardService.Save(dto1)
}

func (r Controller) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.BoardService.GetAll()
		c.JSON(http.StatusOK, gin.H{
			"data": arr,
		})
	}
}

func (r Controller) GetById(id int64) {
	fmt.Println(r.BoardService.GetById(id))
}
