package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
)

type ReplyServiceImpl struct {
	Repository      repository.ReplyRepository
	BoardRepository repository.BoardRepository
}

func NewReplyService(replyRepository repository.ReplyRepository, boardRepository repository.BoardRepository) ReplyServiceImpl {
	return ReplyServiceImpl{replyRepository, boardRepository}
}

func (r ReplyServiceImpl) GetById(id int64) entity.Reply {
	return r.Repository.FindById(id)
}

func (r ReplyServiceImpl) Save(boardId int64, dto1 dto.SaveReplyDto) int64 {
	board := r.BoardRepository.FindById(boardId)

	var arr []entity.Reply
	arr = r.Repository.FindAll()

	var max = int64(0)
	for _, v := range arr {
		if v.Id > max {
			max = v.Id
		}
	}

	id := r.Repository.Save(entity.Reply{
		Id:      max + 1,
		Board:   board,
		Content: dto1.Contents,
	})
	return id
}

func (r ReplyServiceImpl) GetAll() []entity.Reply {
	return r.Repository.FindAll()
}
