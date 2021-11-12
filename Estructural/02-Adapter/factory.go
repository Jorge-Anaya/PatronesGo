package _2_Adapter

func Factory(transporte string) Adapter {
	switch transporte {
	case "automovil":
		elTransporte := Automovil{}
		return &AutomovilAdapter{&elTransporte}
	case "bicicleta":
		elTransporte := Bicicleta{}
		return &BicicletaAdapter{&elTransporte}
	}
	return nil
}