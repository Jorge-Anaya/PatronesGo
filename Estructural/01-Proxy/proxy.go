package _1_Proxy

type Proxy interface {
	GetByID(ID uint) Book
	GetAll() Books
}
