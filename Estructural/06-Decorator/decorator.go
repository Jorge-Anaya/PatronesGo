package _6_Decorator

// Decorator interface que deben implementar las estructuras
// que se van a decorar.
type Decorator interface {
	Process() error
}