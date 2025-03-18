package domain

import (
	"testing"
)

func TestNewMoney(t *testing.T) {
	// Test valid money creation
	amount := 10.5
	currency := "EUR"
	money := NewMoney(amount, currency)
	
	if money == nil {
		t.Errorf("Expected money not to be nil")
	}
	
	if money.Amount() != amount {
		t.Errorf("Expected amount %f, got %f", amount, money.Amount())
	}
	
	if money.Currency() != currency {
		t.Errorf("Expected currency %s, got %s", currency, money.Currency())
	}
}

func TestCurrency(t *testing.T) {
	money := NewMoney(10.0, "USD")
	if money.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", money.Currency())
	}
}

func TestMul(t *testing.T) {
	initialAmount := 10.0
	multiplier := 2.5
	expectedAmount := initialAmount * multiplier
	
	money := NewMoney(initialAmount, "EUR")
	result := money.Mul(multiplier)
	
	if result.Amount() != expectedAmount {
		t.Errorf("Expected amount %f after multiplication, got %f", expectedAmount, result.Amount())
	}
	
	if result.Currency() != "EUR" {
		t.Errorf("Expected currency EUR after multiplication, got %s", result.Currency())
	}
}

func TestAddMoney(t *testing.T) {
	// Test adding money with same currency
	money1 := NewMoney(10.0, "USD")
	money2 := NewMoney(5.0, "USD")
	
	result, err := money1.AddMoney(money2)
	
	if err != nil {
		t.Errorf("Unexpected error when adding money with same currency: %v", err)
	}
	
	expectedAmount := 15.0
	if result.Amount() != expectedAmount {
		t.Errorf("Expected amount %f after addition, got %f", expectedAmount, result.Amount())
	}
	
	// Test adding money with different currency (should return error)
	money3 := NewMoney(10.0, "USD")
	money4 := NewMoney(5.0, "EUR")
	
	_, err = money3.AddMoney(money4)
	
	if err == nil {
		t.Errorf("Expected error when adding money with different currency, got none")
	}
	
	if err != ErrorMoneyCurrency {
		t.Errorf("Expected ErrorMoneyCurrency, got %v", err)
	}
}

func TestDiv(t *testing.T) {
	initialAmount := 10.0
	divisor := 2.0
	expectedAmount := initialAmount / divisor
	
	money := NewMoney(initialAmount, "EUR")
	result := money.Div(divisor)
	
	if result.Amount() != expectedAmount {
		t.Errorf("Expected amount %f after division, got %f", expectedAmount, result.Amount())
	}
	
	if result.Currency() != "EUR" {
		t.Errorf("Expected currency EUR after division, got %s", result.Currency())
	}
}

func TestCents(t *testing.T) {
	testCases := []struct {
		amount       float64
		expectedCents int
	}{
		{1.0, 100},
		{1.5, 150},
		{0.01, 1},
		{10.99, 1099},
	}
	
	for _, tc := range testCases {
		money := NewMoney(tc.amount, "EUR")
		cents := money.Cents()
		
		if cents != tc.expectedCents {
			t.Errorf("For amount %f, expected %d cents, got %d", tc.amount, tc.expectedCents, cents)
		}
	}
}

func TestAmount(t *testing.T) {
	amount := 10.5
	money := NewMoney(amount, "EUR")
	
	if money.Amount() != amount {
		t.Errorf("Expected amount %f, got %f", amount, money.Amount())
	}
}

func TestMoneyString(t *testing.T) {
	money := NewMoney(10.5, "EUR")
	result := money.String()
	
	expectedStr := "10.50 EUR"
	if result != expectedStr {
		t.Errorf("Expected string representation %s, got %s", expectedStr, result)
	}
	
	// Test nil money
	var nilMoney *Money = nil
	nilResult := nilMoney.String()
	
	if nilResult != "<nil>" {
		t.Errorf("Expected string representation of nil money to be '<nil>', got %s", nilResult)
	}
}