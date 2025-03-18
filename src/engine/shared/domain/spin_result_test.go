package domain

import (
	"strings"
	"testing"
)

func TestSpinResultString(t *testing.T) {
	// Create a non-winning spin result
	noWinResult := &SpinResult{
		VisibleReels: nil,
		WinningLines: nil,
		WinAmount:    nil,
		IsWin:        false,
	}
	
	noWinStr := noWinResult.String()
	if !strings.Contains(noWinStr, "No win") {
		t.Errorf("Expected non-winning result to contain 'No win', got %s", noWinStr)
	}
	
	// Create a winning spin result
	visibleReels := NewVisibleReels()
	visibleReels.AddReelResult([]string{"A", "B", "C"})
	visibleReels.AddReelResult([]string{"A", "D", "E"})
	visibleReels.AddReelResult([]string{"A", "F", "G"})
	
	winningLine := &WinningLine{
		LineSymbols: []string{"A", "A", "A"},
		WinSymbol:   "A",
		LinePos:     []int{0, 0, 0},
	}
	
	winResult := &SpinResult{
		VisibleReels: visibleReels,
		WinningLines: []*WinningLine{winningLine},
		WinAmount:    NewMoney(10.0, "EUR"),
		IsWin:        true,
	}
	
	winStr := winResult.String()
	if !strings.Contains(winStr, "WIN!") {
		t.Errorf("Expected winning result to contain 'WIN!', got %s", winStr)
	}
	
	if !strings.Contains(winStr, "10.00 EUR") {
		t.Errorf("Expected winning result to contain winning amount '10.00 EUR', got %s", winStr)
	}
	
	if !strings.Contains(winStr, "Winning Lines") {
		t.Errorf("Expected winning result to contain 'Winning Lines', got %s", winStr)
	}
}