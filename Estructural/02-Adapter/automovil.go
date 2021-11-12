package _2_Adapter

import "fmt"

type Automovil struct {
	Marca     string
	Model     uint8
	Encendido bool
}

func (myAtomovil *Automovil) Encender() {
	if myAtomovil.Encendido {
		fmt.Println("Ya está encendido")
		return
	}
	fmt.Println("Encendido!")
}

func (myAtomovil *Automovil) Acelerar() {
	fmt.Println("Acelerando!")
}