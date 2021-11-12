package _4_Bridge

// ToDo interface que tiene la lista de tareas
type ToDo interface {
	SetList(l List)
	Add(name string)
	Print()
}