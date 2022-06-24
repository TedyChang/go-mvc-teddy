package repository

import (
	"codetest/board/entity"
	"codetest/database"
	"gorm.io/gorm"
)

type BoardModel interface {
	First(dest interface{}, conds ...interface{}) BoardModel
	Find(dest interface{}, conds ...interface{}) BoardModel
	Create(value interface{}) BoardModel
	Where(query interface{}, args ...interface{}) BoardModel
}

type BoardQuery struct {
	*gorm.DB
}

func (r BoardQuery) First(dest interface{}, conds ...interface{}) BoardModel {
	return BoardModel(BoardQuery{r.DB.First(dest, conds)})
}
func (r BoardQuery) Find(dest interface{}, conds ...interface{}) BoardModel {
	return BoardModel(BoardQuery{r.DB.Find(dest, conds)})
}
func (r BoardQuery) Create(value interface{}) BoardModel {
	return BoardModel(BoardQuery{r.DB.Create(value)})
}
func (r BoardQuery) Where(query interface{}, args ...interface{}) BoardModel {
	return BoardModel(BoardQuery{r.DB.Where(query, args)})
}

type BoardRepository struct {
	BoardModel
}

func NewRepository() BoardRepository {
	return BoardRepository{BoardQuery{database.Db}}
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
	i := r.Where("id = ?", id)
	i.First(&board)

	return board.ID != 0
}
