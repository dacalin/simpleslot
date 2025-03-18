package domain

import (
	"testing"
)

func TestErrorConstants(t *testing.T) {
	// Test that error constants are defined
	if ErrorRng == nil {
		t.Errorf("Expected ErrorRng to be defined")
	}
	
	if ErrorMoneyCurrency == nil {
		t.Errorf("Expected ErrorMoneyCurrency to be defined")
	}
	
	if ErrorMoneyAmount == nil {
		t.Errorf("Expected ErrorMoneyAmount to be defined")
	}
	
	if ErrorVisibleReel == nil {
		t.Errorf("Expected ErrorVisibleReel to be defined")
	}
	
	// Test error messages
	if ErrorRng.Error() != "RNG Error" {
		t.Errorf("Expected ErrorRng message to be 'RNG Error', got '%s'", ErrorRng.Error())
	}
	
	if ErrorMoneyCurrency.Error() != "Money::CurrencyError" {
		t.Errorf("Expected ErrorMoneyCurrency message to be 'Money::CurrencyError', got '%s'", ErrorMoneyCurrency.Error())
	}
	
	if ErrorMoneyAmount.Error() != "Money::AmountError" {
		t.Errorf("Expected ErrorMoneyAmount message to be 'Money::AmountError', got '%s'", ErrorMoneyAmount.Error())
	}
	
	if ErrorVisibleReel.Error() != "VisibleReel::Error" {
		t.Errorf("Expected ErrorVisibleReel message to be 'VisibleReel::Error', got '%s'", ErrorVisibleReel.Error())
	}
}