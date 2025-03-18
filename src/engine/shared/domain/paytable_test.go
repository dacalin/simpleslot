package domain

import (
	"testing"
)

func TestNewPayTable(t *testing.T) {
	payTable := NewPayTable()
	
	if payTable == nil {
		t.Errorf("Expected payTable not to be nil")
	}
	
	// Check that the multipliers map was initialized
	if payTable.multipliers == nil {
		t.Errorf("Expected multipliers map to be initialized")
	}
}

func TestPayTableAdd(t *testing.T) {
	payTable := NewPayTable()
	
	// Add a symbol with a multiplier
	symbol := "A"
	multiplier := 10
	payTable.Add(symbol, multiplier)
	
	// Verify the symbol was added with correct multiplier
	if payTable.GetMultiplier(symbol) != multiplier {
		t.Errorf("Expected multiplier %d for symbol %s, got %d", multiplier, symbol, payTable.GetMultiplier(symbol))
	}
}

func TestGetBestSymbol(t *testing.T) {
	payTable := NewPayTable()
	
	// Add multiple symbols with different multipliers
	payTable.Add("A", 10)
	payTable.Add("B", 5)
	payTable.Add("C", 15)
	payTable.Add("D", 8)
	
	// Get the best symbol (should be "C" with multiplier 15)
	bestSymbol := payTable.GetBestSymbol()
	
	if bestSymbol != "C" {
		t.Errorf("Expected best symbol to be \"C\", got %s", bestSymbol)
	}
}

func TestEmptyPayTable(t *testing.T) {
	payTable := NewPayTable()
	
	// Getting a multiplier for a non-existent symbol should return 0
	multiplier := payTable.GetMultiplier("X")
	if multiplier != 0 {
		t.Errorf("Expected multiplier 0 for non-existent symbol, got %d", multiplier)
	}
	
	// Getting best symbol from empty paytable should return empty string
	bestSymbol := payTable.GetBestSymbol()
	if bestSymbol != "" {
		t.Errorf("Expected empty string for best symbol in empty paytable, got %s", bestSymbol)
	}
}