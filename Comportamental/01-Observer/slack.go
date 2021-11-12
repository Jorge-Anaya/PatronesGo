package _1_Observer

import (
	"fmt"
)

// Slack .
type Slack struct{}

// Notify .
func (s *Slack) Notify(data string) {
	sendMessage(data)
}

func sendMessage(data string) {
	fmt.Printf("He enviado un slack con el mensaje: %q\n", data)
}