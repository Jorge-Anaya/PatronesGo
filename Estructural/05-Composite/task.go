package _5_Composite

import (
	"strconv"
	"strings"
)

// Task ...
type Task struct {
	name        string
	responsable string
	price       int
	subTasks    []Item
}

// Add ...
func (myTask *Task) Add(myItem Item) {
	myTask.subTasks = append(myTask.subTasks, myItem)
}

// String ...
func (myTask *Task) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t|--")
	sb.WriteString(myTask.name)
	sb.WriteString(" - ")
	sb.WriteString(myTask.responsable)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(myTask.Price()))
	sb.WriteString("\n")
	for _, v := range myTask.subTasks {
		sb.WriteString(v.String())
	}
	return sb.String()
}

// Price ...
func (myTask *Task) Price() int {
	price := myTask.price
	for _, v := range myTask.subTasks {
		price += v.Price()
	}
	return price
}