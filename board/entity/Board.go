package entity

type Rows struct {
	BoardRows []Board
	ReplyRows []Reply
}

type Board struct {
	Id       int64
	Title    string
	Contents string
}
