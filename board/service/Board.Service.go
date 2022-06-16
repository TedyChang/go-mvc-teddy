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
	var arr []entity.Board
	arr = r.Repository.FindAll()

	var max = int64(0)
	for _, v := range arr {
		if v.Id > max {
			max = v.Id
		}
	}

	id := r.Repository.Save(entity.Board{
		Id:       max + 1,
		Title:    dto1.Title,
		Contents: dto1.Contents,
	})
	return id
}

func (r BoardService) GetAll() []entity.Board {
	return r.Repository.FindAll()
}
