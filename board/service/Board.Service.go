package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
)

type BoardServiceImpl struct {
	Repository repository.BoardRepository
}

func NewService(boardRepository repository.BoardRepository) BoardServiceImpl {
	return BoardServiceImpl{boardRepository}
}

func (r BoardServiceImpl) GetById(id int64) entity.Board {
	return r.Repository.FindById(id)
}

func (r BoardServiceImpl) Save(dto1 dto.SaveDto) int64 {
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

func (r BoardServiceImpl) GetAll() []entity.Board {
	return r.Repository.FindAll()
}
