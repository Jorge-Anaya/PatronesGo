package _6_Decorator

import "fmt"

type Route struct {
	decorators map[string]Decorator
}

func NewRoute() Route {
	return Route{make(map[string]Decorator)}
}

func (myRoute *Route) Add(decorator Decorator, path string) {
	myRoute.decorators[path] = decorator
}

func (myRoute *Route) Exec(path string) {
	d, ok := myRoute.decorators[path]
	if !ok {
		fmt.Println("no existe esa ruta")
		return
	}
	err := d.Process()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}