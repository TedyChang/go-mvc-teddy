package controller

import (
	"codetest/board/dto"
	"codetest/board/service"
	"codetest/util/gu"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReplyController struct {
	ReplyService service.ReplyService
}

func NewReplyController(impl service.ReplyService) ReplyController {
	return ReplyController{impl}
}

func (r ReplyController) SaveReply() func(c *gin.Context) {
	return func(c *gin.Context) {
		boardId, err1 := gu.GetID("board", c)
		if err1 != nil {
			gu.BadJson(c, gin.H{"message": "error : board is not number"})
			return
		}

		var replyDto dto.SaveReplyDto

		err2 := c.ShouldBindJSON(&replyDto)
		if err2 != nil {
			gu.BadJson(c, gin.H{"message": "error : not SaveBoard"})
			return
		}

		gu.OkJson(c, gin.H{"id": r.ReplyService.Save(boardId, replyDto)})
	}
}

func (r ReplyController) GetAllReply() func(c *gin.Context) {
	return func(c *gin.Context) {
		arr := r.ReplyService.GetAll()

		gu.OkJson(c, gin.H{"data": arr})
	}
}

func (r ReplyController) GetReplyById() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err1 := gu.GetID("reply", c)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error : reply is not number"})
			return
		}

		gu.OkJson(c, gin.H{"data": r.ReplyService.GetById(id)})
	}
}
