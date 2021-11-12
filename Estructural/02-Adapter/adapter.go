package _2_Adapter

type Adapter interface {
	Mover()
}

type AutomovilAdapter struct {
	Auto *Automovil
}

func (myAutomovil *AutomovilAdapter) Mover() {
	if !myAutomovil.Auto.Encendido {
		myAutomovil.Auto.Encender()
	}
	myAutomovil.Auto.Acelerar()
}

type BicicletaAdapter struct {
	Bici *Bicicleta
}

func (myBicicleta *BicicletaAdapter) Mover() {
	myBicicleta.Bici.Avanzar()
}
