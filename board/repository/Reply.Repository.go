package repository

import (
	"codetest/board/entity"
	"codetest/database"
	"gorm.io/gorm"
)

type ReplyModel interface {
	First(dest interface{}, conds ...interface{}) ReplyModel
	Find(dest interface{}, conds ...interface{}) ReplyModel
	Create(value interface{}) ReplyModel
	Where(query interface{}, args ...interface{}) ReplyModel
}

type ReplyQuery struct {
	*gorm.DB
}

func (r ReplyQuery) First(dest interface{}, conds ...interface{}) ReplyModel {
	return ReplyQuery{r.DB.First(dest, conds)}
}
func (r ReplyQuery) Find(dest interface{}, conds ...interface{}) ReplyModel {
	return ReplyQuery{r.DB.Find(dest, conds)}
}
func (r ReplyQuery) Create(value interface{}) ReplyModel {
	return ReplyQuery{r.DB.Create(value)}
}
func (r ReplyQuery) Where(query interface{}, args ...interface{}) ReplyModel {
	return ReplyQuery{r.DB.Where(query, args)}
}

type ReplyRepository struct {
	ReplyModel
}

func NewReplyRepository() ReplyRepository {
	return ReplyRepository{ReplyQuery{database.Db}}
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
