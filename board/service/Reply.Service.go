package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
)

type ReplyService struct {
	Repository      repository.ReplyRepository
	BoardRepository repository.BoardRepository
}

func NewReplyService(replyRepository repository.ReplyRepository, boardRepository repository.BoardRepository) ReplyService {
	return ReplyService{replyRepository, boardRepository}
}

func (r ReplyService) GetById(id int64) entity.Reply {
	return r.Repository.FindById(id)
}

func (r ReplyService) Save(boardId int64, dto1 dto.SaveReplyDto) int64 {
	isExistBoard := r.BoardRepository.IsExist(boardId)
	if !isExistBoard {
		return 0
	}

	id := r.Repository.Save(entity.Reply{
		Board:   entity.Board{ID: boardId},
		Content: dto1.Contents,
	})
	return id
}

func (r ReplyService) GetAll() []entity.Reply {
	return r.Repository.FindAll()
}
