package dto

type SaveBoardDto struct {
	Id       int64
	Title    string
	Contents string
}

func (r *SaveBoardDto) Setter(id int64, title string, contents string) {
	r.Id = id
	r.Title = title
	r.Contents = contents
}

func (r *SaveBoardDto) Getter(id int64, title string, contents string) {
	r.Id = id
	r.Title = title
	r.Contents = contents
}
