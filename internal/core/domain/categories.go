package domain

type Category struct {
	ID          int
	ProductId   []int
	Name        string
	Description string // can be null
}
