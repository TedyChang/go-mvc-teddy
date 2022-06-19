package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
)

type BoardService struct {
	Repository repository.BoardRepository
}

func NewService(boardRepository repository.BoardRepository) BoardService {
	return BoardService{boardRepository}
}

func (r BoardService) GetById(id int64) entity.Board {
	return r.Repository.FindById(id)
}

func (r BoardService) Save(dto1 dto.SaveBoardDto) int64 {
	id := r.Repository.Save(entity.Board{
		Title:    dto1.Title,
		Contents: dto1.Contents,
	})
	return id
}

func (r BoardService) GetAll() []entity.Board {
	return r.Repository.FindAll()
}
