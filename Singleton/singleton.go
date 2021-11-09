package Singleton

var personInstance *person

func getInstance() *person{
	if personInstance == nil{
		personInstance = &person{}
	}
	return personInstance
}

type person struct{
	name string
	age int
}

func (person *person) setName(name string){
	person.name = name
}

func (person *person) getName() string{
	return person.name
}

func (person *person) incrementAge(){
	person.age++
}
