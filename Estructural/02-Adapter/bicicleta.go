package _2_Adapter

import "fmt"

type Bicicleta struct {
	Marca string
	Color string
}

func (myBicicleta *Bicicleta) Avanzar() {
	fmt.Println("Avanzando!")
}