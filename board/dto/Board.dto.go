package dto

type SaveBoardDto struct {
	Title    string `json:"title" binding:"required"`
	Contents string `json:"contents"`
}
