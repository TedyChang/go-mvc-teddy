package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"fmt"
)

type Controller struct {
	BoardService boardService.BoardService
}

func NewController(impl boardService.BoardService) Controller {
	return Controller{impl}
}

func (c Controller) SaveBoard(title string, contents string) {
	dto1 := dto.SaveBoardDto{
		Title:    title,
		Contents: contents,
	}

	c.BoardService.Save(dto1)
}

func (c Controller) GetAll() {
	arr := c.BoardService.GetAll()

	for _, value := range arr {
		fmt.Println(value)
	}
}

func (c Controller) GetById(id int64) {
	fmt.Println(c.BoardService.GetById(id))
}
