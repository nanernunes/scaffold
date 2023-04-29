package scaffold

type Controller interface {
	Index(*Context)
	Show(*Context)
	Create(*Context)
	Update(*Context)
	Delete(*Context)
}
