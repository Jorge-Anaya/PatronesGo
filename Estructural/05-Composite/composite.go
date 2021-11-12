package _5_Composite

// Item ...
type Item interface {
	Add(Item)
	String() string
	Price() int
}