package _5_Composite

import (
	"strconv"
	"strings"
)

// Project ...
type Project struct {
	name  string
	tasks []Item
}

// Add ...
func (myProject *Project) Add(item Item) {
	myProject.tasks = append(myProject.tasks, item)
}

// String ...
func (myProject *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(myProject.name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(myProject.Price()))
	sb.WriteString("\n")
	for _, v := range myProject.tasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

// Price ...
func (myProject *Project) Price() int {
	price := 0
	for _, v := range myProject.tasks {
		price += v.Price()
	}
	return price
}
