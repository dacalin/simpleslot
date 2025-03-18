package domain

type Symbol struct {
	value      string
}


func NewSymbol(value string) *Symbol {
	return &Symbol{value: value}
}


