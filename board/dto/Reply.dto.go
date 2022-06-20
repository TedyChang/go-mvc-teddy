package dto

type SaveReplyDto struct {
	Contents string `json:"contents" binding:"required"`
}
