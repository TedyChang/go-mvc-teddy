package main

import (
	"codetest/board/dto"
	boardRepository "codetest/board/repository"
	boardService "codetest/board/service"
	"codetest/database"
	"codetest/layer"
	userRepository "codetest/user/repository"
	userService "codetest/user/service"
)

func LoadDI() (layer.Service, layer.Service) {
	boardRepository := boardRepository.BoardRepository{Model: &database.Db.TBoard}
	boardServiceImpl := boardService.BoardServiceImpl{
		Repository: boardRepository,
	}
	var boardServiceDI layer.Service = boardServiceImpl

	userRepository := userRepository.UserRepository{Model: &database.Db.TUser}
	userServiceImpl := userService.UserServiceImpl{
		Repository: userRepository,
	}
	var userServiceDI layer.Service = userServiceImpl

	return boardServiceDI, userServiceDI
}

func main() {
	database.Db.DbSetting()

	// mvc struct, interface & DI
	boardServiceDI, userServiceDI := LoadDI()

	// test boardService
	// session 001 board service di

	boardServiceDI.GetById(1)
	dto1 := dto.SaveDto{
		Id:       4,
		Title:    "추가된 글입니다",
		Contents: "추가된 내용입니다",
	}
	dto2 := dto.SaveDto{
		Id:       5,
		Title:    "오늘은 날씨가 좋내요",
		Contents: "밖에서 밥먹기 좋은 날씨입니다.",
	}
	dto3 := dto.SaveDto{
		Id:       6,
		Title:    "졸릴땐 커피죠",
		Contents: "회사 앞 굿커피 한잔하고 오겠습니다 :)",
	}
	boardServiceDI.(boardService.BoardServiceImpl).Save(dto1)
	boardServiceDI.(boardService.BoardServiceImpl).Save(dto2)
	boardServiceDI.(boardService.BoardServiceImpl).Save(dto3)
	boardServiceDI.GetById(4)
	boardServiceDI.(boardService.BoardServiceImpl).GetAll()

	// session 002 user service di

	userServiceDI.GetById(1)
	userServiceDI.(userService.UserServiceImpl).Save(7, "저스틴")
	userServiceDI.(userService.UserServiceImpl).GetAll()

}
