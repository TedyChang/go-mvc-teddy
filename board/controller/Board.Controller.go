package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"fmt"
)

type Controller struct {
	BoardService boardService.BoardServiceImpl
}

func NewController(impl boardService.BoardServiceImpl) Controller {
	return Controller{impl}
}

func (c Controller) SaveBoard(id int64, title string, contents string) {
	dto1 := dto.SaveBoardDto{
		Id:       id,
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
