package dto

type SaveDto struct {
	Id       int64
	Title    string
	Contents string
}

func (r *SaveDto) Setter(id int64, title string, contents string) {
	r.Id = id
	r.Title = title
	r.Contents = contents
}

func (r *SaveDto) Getter(id int64, title string, contents string) {
	r.Id = id
	r.Title = title
	r.Contents = contents
}
