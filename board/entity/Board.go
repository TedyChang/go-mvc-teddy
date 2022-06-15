package entity

type Rows struct {
	Rows []Board
}

type Board struct {
	Id       int64
	Title    string
	Contents string
}
