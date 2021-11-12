package _4_Bridge

// Any recibe tareas repetidas
type Any struct {
	rendering List
	todos     []string
}

func NewAny() *Any {
	return &Any{}
}

func (a *Any) SetList(l List) {
	a.rendering = l
}

// Add agrega una tarea a la lista
func (a *Any) Add(name string) {
	a.todos = append(a.todos, name)
}

// Print muestra la lista con su
// respectiva representación
func (a *Any) Print() {
	a.rendering.Print(a.todos)
}
