package _4_Bridge

func TestBridge() {
	myToDos := factoryToDo("any")
	rendering := factoryList("bullet")
	myToDos.SetList(rendering)

	myToDos.Add("Qué estudiar?")
	myToDos.Add("Explicarlo con palabras sencillas")
	myToDos.Add("Hacer con lo que aprendimos")
	myToDos.Add("Revisar y mejorar")
	myToDos.Add("Qué estudiar?")
	myToDos.Print()
}

func factoryToDo(s string) ToDo {
	if s == "unique" {
		return NewUnique()
	}
	return NewAny()
}

func factoryList(s string) List {
	if s == "plain" {
		return NewPlain()
	}
	return NewBullet('*')
}