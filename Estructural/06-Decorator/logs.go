package _6_Decorator

import "fmt"

type LogRegistry struct {
	Handler Decorator
}

func NewLogRegistry(handler Decorator) *LogRegistry {
	return &LogRegistry{handler}
}

func (myLogRegistry *LogRegistry) Process() error {
	defer fmt.Println("Petición")
	return myLogRegistry.Handler.Process()
}