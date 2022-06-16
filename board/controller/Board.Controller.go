package controller

import (
	"codetest/board/dto"
	boardService "codetest/board/service"
	"codetest/layer"
)

type Controller struct {
	BoardService layer.Service
	//replyRepository layer.Service
}

func (c Controller) SaveBoard(id int64, title string, contents string) {
	dto1 := dto.SaveDto{
		Id:       id,
		Title:    title,
		Contents: contents,
	}

	c.BoardService.(boardService.BoardServiceImpl).Save(dto1)
}

func (c Controller) GetAll() {
	c.BoardService.(boardService.BoardServiceImpl).GetAll()
}

func (c Controller) GetById(id int64) {
	c.BoardService.(boardService.BoardServiceImpl).GetById(id)
}
