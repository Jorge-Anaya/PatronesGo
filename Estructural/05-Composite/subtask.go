package _5_Composite

import (
	"fmt"
	"strings"
)

// SubTask ...
type SubTask struct {
	name  string
	price int
}

// Add ...
func (mySubTask *SubTask) Add(i Item) {
	fmt.Println("Yo no acepto más subtareas")
}

// String ...
func (mySubTask *SubTask) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t\t|-- ")
	sb.WriteString(mySubTask.name)
	sb.WriteString("\n")

	return sb.String()
}

// Price ...
func (mySubTask *SubTask) Price() int {
	return mySubTask.price
}