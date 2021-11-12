package _1_Observer

type Observable interface {
	AddObserver(name string, o Observer)
	RemoveObserver(name string)
	NotifyObservers()
}