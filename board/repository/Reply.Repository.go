package repository

import (
	"codetest/board/entity"
	"codetest/database"
	"fmt"
)

type ReplyRepositoryImpl struct {
	Model *entity.Rows
}

type ReplyRepository interface {
	FindById(id int64) entity.Reply
	FindAll() []entity.Reply
	Save(reply entity.Reply) int64
}

func NewReplyRepository() ReplyRepository {
	return ReplyRepositoryImpl{&database.Db.TBoard}
}

func (r ReplyRepositoryImpl) FindById(id int64) entity.Reply {
	result := entity.Reply{}
	for _, reply := range r.Model.ReplyRows {
		if isReplyId(reply, id) {
			result = reply
			break
		}
	}
	return result
}

func (r ReplyRepositoryImpl) FindAll() []entity.Reply {
	arr := make([]entity.Reply, len(r.Model.ReplyRows))

	for i, reply := range r.Model.ReplyRows {
		arr[i] = reply
	}
	return arr
}

func (r ReplyRepositoryImpl) Save(reply entity.Reply) int64 {
	r.Model.ReplyRows = append(r.Model.ReplyRows, reply)
	fmt.Println("save : ", reply.Content)
	return reply.Id
}

func isReplyId(reply entity.Reply, id int64) bool {
	return reply.Id == id
}
