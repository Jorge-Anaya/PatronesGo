package _1_Proxy

type Local struct {
	Remote *Data
	cache  Books
}

func NewProxy() Proxy {
	return &Local{
		Remote: NewBD("https://alguno.com", 8080, "usuario", "123456"),
		cache:  make(Books, 0),
	}
}

func (l *Local) GetByID(ID uint) Book {
	for _, v := range l.cache {
		if v.ID == ID {
			return v
		}
	}

	b := l.Remote.ByID(ID)
	l.cache = append(l.cache, b)

	return b
}

func (l *Local) GetAll() Books {
	return l.Remote.All()
}
