package _4_Prototype

import "fmt"

type prototype struct {
	name    string
	age     int
	friends []string
	color   *string
	phones  map[string]string
}

func (myPrototype *prototype) String() string {
	return fmt.Sprintf(
		"Nombre: %s, Edad: %d, Amigos: %v, Color: %s, Teléfonos: %v",
		myPrototype.name,
		myPrototype.age,
		myPrototype.friends,
		*myPrototype.color,
		myPrototype.phones,
	)
}

func (myPrototype *prototype) Clone() prototype {
	friends := make([]string, len(myPrototype.friends), len(myPrototype.friends))
	copy(friends, myPrototype.friends)

	color := *myPrototype.color

	phones := make(map[string]string)
	for k, v := range myPrototype.phones {
		phones[k] = v
	}

	return prototype{
		name:    myPrototype.name,
		age:     myPrototype.age,
		friends: friends,
		color:   &color,
		phones:  phones,
	}
}
