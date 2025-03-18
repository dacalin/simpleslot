package domain

import (
	"testing"
)

func TestNewSymbol(t *testing.T) {
	// Test creating a new symbol
	symbolValue := "A"
	symbol := NewSymbol(symbolValue)
	
	// Check that the symbol was created correctly
	if symbol == nil {
		t.Errorf("Expected symbol not to be nil")
	}
}