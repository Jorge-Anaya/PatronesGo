package _1_Singleton

import "sync"

var (
	personInstance *person
	once sync.Once
)

func GetInstance() *person {
	once.Do(func() {
		personInstance = &person{}
	})
	return personInstance
}

type person struct{
	name string
	age int
	mux sync.Mutex
}

func (person *person) SetName(name string) {
	person.mux.Lock()
	defer person.mux.Unlock()
	person.name = name
}

func (person *person) GetName() string {
	person.mux.Lock()
	defer person.mux.Unlock()
	return person.name
}

func (person *person) SetAge(age int) {
	person.mux.Lock()
	defer person.mux.Unlock()
	person.age = age
}

func (person *person) GetAge() int {
	person.mux.Lock()
	defer person.mux.Unlock()
	return person.age
}

func (person *person) IncrementAge()  {
	person.mux.Lock()
	defer person.mux.Unlock()
	person.age++
}

