package repository

import (
	"codetest/board/entity"
	"codetest/database"
	"gorm.io/gorm"
)

type Model interface {
	Model(interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
}

type BoardRepository struct {
	Model
}

func NewRepository() BoardRepository {
	return BoardRepository{database.Db}
}

func (r BoardRepository) FindById(id int64) entity.Board {
	b := entity.Board{}
	r.First(&b, "id = ? ", id)
	return b
}

func (r BoardRepository) FindAll() []entity.Board {
	var boards []entity.Board
	r.Find(&boards)
	return boards
}

func (r BoardRepository) Save(board entity.Board) int64 {
	_ = r.Create(&board)
	return board.ID
}

func (r BoardRepository) IsExist(id int64) bool {
	var board entity.Board
	r.Where("id = ?", id).First(&board)

	return board.ID != 0
}
