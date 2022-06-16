package main

import (
	boardController "codetest/board/controller"
	boardRepository "codetest/board/repository"
	boardService "codetest/board/service"
	"codetest/database"
	"codetest/layer"
	userRepository "codetest/user/repository"
	userService "codetest/user/service"
)

func LoadDI() (layer.Controller, layer.Service) {
	var boardRepositoryConfig = boardRepository.BoardRepositoryImpl{Model: &database.Db.TBoard}
	boardServiceImpl := boardService.BoardServiceImpl{
		Repository: boardRepositoryConfig,
	}

	var boardServiceDI layer.Service = boardServiceImpl

	userRepositoryConfig := userRepository.UserRepository{Model: &database.Db.TUser}
	userServiceImpl := userService.UserServiceImpl{
		Repository: userRepositoryConfig,
	}
	var userServiceDI layer.Service = userServiceImpl

	var boardControllerDI = boardController.Controller{
		BoardService: boardServiceDI,
	}

	return boardControllerDI, userServiceDI
}

func main() {
	database.Db.DbSetting()

	// mvc struct, interface & DI
	boardControllerDI, userServiceDI := LoadDI()

	// test boardService
	// session 001 board service di

	boardControllerDI.(boardController.Controller).SaveBoard(7, "컨트롤러에서 올린 글입니다.", "내용은 없습니다")
	boardControllerDI.(boardController.Controller).GetAll()

	// session 002 user service di

	userServiceDI.GetById(1)
	userServiceDI.(userService.UserServiceImpl).Save(7, "저스틴")
	userServiceDI.(userService.UserServiceImpl).GetAll()

}
