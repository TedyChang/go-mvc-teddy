package layer

type Repository interface {
	FindById(int64)
	//Save(any2 any)
}

type Service interface {
	GetById(int64)
}

type Controller interface {
}
