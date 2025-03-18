package domain

import (
	"testing"
)

func TestNewReel(t *testing.T) {
	id := 1
	reel := NewReel(id)
	
	if reel == nil {
		t.Errorf("Expected reel not to be nil")
	}
}

func TestReelAdd(t *testing.T) {
	reel := NewReel(1)
	
	// Add symbols with biases
	reel = reel.Add("A", 10)
	reel = reel.Add("B", 5)
	reel = reel.Add("C", 15)
	
	// Verify total bias is correct (10 + 5 + 15 = 30)
	expectedTotalBias := 30
	if reel.TotalBias() != expectedTotalBias {
		t.Errorf("Expected total bias %d, got %d", expectedTotalBias, reel.TotalBias())
	}
}

func TestGetSymbolFromCumulativeBias(t *testing.T) {
	reel := NewReel(1)
	
	// Add symbols with biases: A(10), B(5), C(15)
	reel = reel.Add("A", 10)
	reel = reel.Add("B", 5)
	reel = reel.Add("C", 15)
	
	testCases := []struct {
		bias          int
		expectedSymbol string
	}{
		{5, "A"},     // Within first symbol's bias
		{10, "A"},    // At edge of first symbol's bias
		{11, "B"},    // Within second symbol's bias
		{15, "B"},    // At edge of second symbol's bias
		{20, "C"},    // Within third symbol's bias
		{30, "C"},    // At edge of third symbol's bias
		// Testing values beyond total bias is implementation dependent
		// May need to adjust this test based on code's behavior
	}
	
	for _, tc := range testCases {
		result := reel.GetSymbolFromCumulativeBias(tc.bias)
		if result != tc.expectedSymbol {
			t.Errorf("For bias %d, expected symbol %s, got %s", tc.bias, tc.expectedSymbol, result)
		}
	}
}