package repository

import (
	"codetest/board/entity"
	"codetest/database"
)

type ReplyRepository struct {
	Model
}

func NewReplyRepository() ReplyRepository {
	return ReplyRepository{database.Db}
}

func (r ReplyRepository) FindById(id int64) entity.Reply {
	reply := entity.Reply{}
	r.First(&reply, "id = ? ", id)
	return reply
}

func (r ReplyRepository) FindAll() []entity.Reply {
	var replies []entity.Reply
	r.Find(&replies)
	return replies
}

func (r ReplyRepository) Save(reply entity.Reply) int64 {
	_ = r.Create(&reply)
	return reply.ID
}
