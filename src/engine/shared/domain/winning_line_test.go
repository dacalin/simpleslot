package domain

import (
	"strings"
	"testing"
)

func TestWinningLineString(t *testing.T) {
	// Test with a valid winning line
	wl := &WinningLine{
		LineSymbols: []string{"A", "A", "A"},
		WinSymbol:   "A",
		LinePos:     []int{0, 0, 0},
	}
	
	result := wl.String()
	
	// Check that the result contains all the symbols
	for _, symbol := range wl.LineSymbols {
		if !strings.Contains(result, symbol) {
			t.Errorf("Expected string representation to contain symbol %s, got %s", symbol, result)
		}
	}
	
	// Test with nil winning line
	var nilWL *WinningLine = nil
	nilResult := nilWL.String()
	
	if nilResult != "<nil>" {
		t.Errorf("Expected string representation of nil winning line to be '<nil>', got %s", nilResult)
	}
}