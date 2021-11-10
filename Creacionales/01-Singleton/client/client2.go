package client

import "github.com/Jorge-Anaya/PatronesGo/Creacionales/01-Singleton"

func IncrementAgeClient2() {
	person:= _1_Singleton.GetInstance()
	person.IncrementAge()
}
