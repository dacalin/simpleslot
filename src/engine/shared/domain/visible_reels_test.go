package domain

import (
	"strings"
	"testing"
)

func TestNewVisibleReels(t *testing.T) {
	vr := NewVisibleReels()
	
	if vr == nil {
		t.Errorf("Expected visible reels not to be nil")
	}
	
	if vr.ReelsCount() != 0 {
		t.Errorf("Expected initial reels count to be 0, got %d", vr.ReelsCount())
	}
}

func TestAddReelResult(t *testing.T) {
	vr := NewVisibleReels()
	
	// Add first reel
	reel1 := []string{"A", "B", "C"}
	vr.AddReelResult(reel1)
	
	if vr.ReelsCount() != 1 {
		t.Errorf("Expected reels count to be 1 after adding first reel, got %d", vr.ReelsCount())
	}
	
	// Add second reel
	reel2 := []string{"D", "E", "F"}
	vr.AddReelResult(reel2)
	
	if vr.ReelsCount() != 2 {
		t.Errorf("Expected reels count to be 2 after adding second reel, got %d", vr.ReelsCount())
	}
}

func TestGetSymbol(t *testing.T) {
	vr := NewVisibleReels()
	
	// Add reels
	reel1 := []string{"A", "B", "C"}
	reel2 := []string{"D", "E", "F"}
	vr.AddReelResult(reel1)
	vr.AddReelResult(reel2)
	
	// Test valid positions
	testCases := []struct {
		reelId        int
		rowId         int
		expectedSymbol string
		expectError   bool
	}{
		{0, 0, "A", false},
		{0, 1, "B", false},
		{0, 2, "C", false},
		{1, 0, "D", false},
		{1, 1, "E", false},
		{1, 2, "F", false},
		{2, 0, "", true},  // Invalid reel
		{0, 3, "", true},  // Invalid row
	}
	
	for _, tc := range testCases {
		symbol, err := vr.Get(tc.reelId, tc.rowId)
		
		if tc.expectError && err == nil {
			t.Errorf("Expected error for reel %d, row %d, but got none", tc.reelId, tc.rowId)
		}
		
		if !tc.expectError && err != nil {
			t.Errorf("Unexpected error for reel %d, row %d: %v", tc.reelId, tc.rowId, err)
		}
		
		if !tc.expectError && symbol != tc.expectedSymbol {
			t.Errorf("For reel %d, row %d: expected symbol %s, got %s", tc.reelId, tc.rowId, tc.expectedSymbol, symbol)
		}
	}
	
	// Note: The Get method doesn't currently check for negative indices
	// If implementation changes to handle negative indices, add tests for that
}

func TestEmptyVisibleReels(t *testing.T) {
	vr := NewVisibleReels()
	
	// Trying to get a symbol from empty reels should return an error
	_, err := vr.Get(0, 0)
	if err == nil {
		t.Errorf("Expected error when getting symbol from empty reels, got none")
	}
	
	if err != ErrorVisibleReel {
		t.Errorf("Expected ErrorVisibleReel, got %v", err)
	}
}

func TestString(t *testing.T) {
	vr := NewVisibleReels()
	
	// Test empty reels
	emptyString := vr.String()
	if !strings.Contains(emptyString, "<nil>") {
		t.Errorf("Expected string representation of empty reels to contain '<nil>', got %s", emptyString)
	}
	
	// Add reels
	vr.AddReelResult([]string{"A", "B", "C"})
	vr.AddReelResult([]string{"D", "E", "F"})
	
	// Test non-empty reels
	nonEmptyString := vr.String()
	if !strings.Contains(nonEmptyString, "A") || !strings.Contains(nonEmptyString, "F") {
		t.Errorf("Expected string representation to contain symbols, got %s", nonEmptyString)
	}
}