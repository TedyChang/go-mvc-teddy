package controller

import (
	"codetest/board/dto"
	"codetest/board/service"
	"fmt"
)

type ReplyController struct {
	ReplyService service.ReplyServiceImpl
}

func NewReplyController(impl service.ReplyServiceImpl) ReplyController {
	return ReplyController{impl}
}

func (c ReplyController) SaveReply(boardId int64, contents string) {
	dto1 := dto.SaveReplyDto{
		Contents: contents,
	}

	c.ReplyService.Save(boardId, dto1)
}

func (c ReplyController) GetAllReply() {
	arr := c.ReplyService.GetAll()

	for _, value := range arr {
		fmt.Println(value)
	}
}

func (c ReplyController) GetReplyById(id int64) {
	fmt.Println(c.ReplyService.GetById(id))
}
