package database

import board "codetest/board/entity"
import user "codetest/user/entity"

var Db DbTable

type DbTable struct {
	TBoard board.Rows
	TUser  user.Rows
}

func (r *DbTable) DbSetting() {
	r.TBoard = board.Rows{
		Rows: []board.Board{
			{Id: 1, Title: "테스트 글입니다 1", Contents: "내용입니다 1"},
			{Id: 2, Title: "테스트 글입니다 2", Contents: "내용입니다 2"},
			{Id: 3, Title: "테스트 글입니다 3", Contents: "내용입니다 3"},
		},
	}
	r.TUser = user.Rows{
		Rows: []user.User{
			{Id: 1, Name: "테디"},
			{Id: 2, Name: "데이빗"},
			{Id: 3, Name: "제이"},
			{Id: 4, Name: "세바스찬"},
			{Id: 5, Name: "호세"},
			{Id: 6, Name: "브라이언"},
		},
	}
}
