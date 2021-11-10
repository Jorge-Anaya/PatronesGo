package client

import "github.com/Jorge-Anaya/PatronesGo/Creacionales/01-Singleton"

func IncrementAgeClient1() {
	person:= _1_Singleton.GetInstance()
	person.IncrementAge()
}
