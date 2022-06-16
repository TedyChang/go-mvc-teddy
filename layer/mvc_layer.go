package layer

type Repository interface {
	FindById(int64)
	//Save(any2 any)
}

type Service interface {
}

type Controller interface {
}
