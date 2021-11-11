package http

type RequestItem interface {
	All(route string)
	FindById(route string)
	Insert(route string)
}

type RequestUser interface {
	Login(route string)
	Insert(route string)
}

type Drive interface {
	Run()
}

type Http struct {
	Drive Drive
	Item  RequestItem
	User  RequestUser
}
