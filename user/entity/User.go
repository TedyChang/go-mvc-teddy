package entity

type Rows struct {
	Rows []User
}

type User struct {
	Id   int64
	Name string
}
