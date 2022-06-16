package repository

import (
	"codetest/board/entity"
	"fmt"
)

type ReplyRepository struct {
	Model *entity.Rows
}

func (r ReplyRepository) FindById(id int64) {
	for _, reply := range r.Model.ReplyRows {
		if isReplyId(reply, id) {
			fmt.Println(reply)
		}
	}
}

func isReplyId(reply entity.Reply, id int64) bool {
	return reply.Id == id
}
