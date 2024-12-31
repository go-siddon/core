package core

// model(db, table, model).find().one().many().limit()

// Database interfaces

type Database interface {
}

type Model[T any] interface {
	Find() Find[T]
	Save() Save[T]
	Update() Update[T]
	Delete() Delete[T]
	Index() Index[T]
}

type Find[T any] interface {
	One()
	Many()
}

type Save[T any] interface {
	One()
	Many()
}

type Update[T any] interface {
	One()
	Many()
}

type Delete[T any] interface {
	One()
	Many()
}

type Index[T any] interface {
	One()
	Many()
}
