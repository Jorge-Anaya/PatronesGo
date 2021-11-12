package _1_Observer

// Message .
type Message struct {
	observers map[string]Observer
	Msg       string
}

// AddObserver .
func (m *Message) AddObserver(name string, o Observer) {
	if m.observers == nil {
		m.observers = make(map[string]Observer)
	}
	m.observers[name] = o
}

// RemoveObserver .
func (m *Message) RemoveObserver(name string) {
	delete(m.observers, name)
}

// NotifyObservers .
func (m *Message) NotifyObservers() {
	for _, v := range m.observers {
		v.Notify(m.Msg)
	}
}