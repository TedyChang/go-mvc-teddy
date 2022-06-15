package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
)

type BoardServiceImpl struct {
	Repository repository.BoardRepository
}

func (r BoardServiceImpl) GetById(id int64) {
	r.Repository.FindById(id)
}

func (r BoardServiceImpl) Save(dto1 dto.SaveDto) {
	r.Repository.Save(entity.Board{
		Id:       dto1.Id,
		Title:    dto1.Title,
		Contents: dto1.Contents,
	})
}

func (r BoardServiceImpl) GetAll() {
	r.Repository.FindAll()
}
