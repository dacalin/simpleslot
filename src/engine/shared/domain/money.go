package domain

import (
	"fmt"
	"math/big"
)

type Money struct {
	amount *big.Float
	currency string
}

func NewMoney(value float64, currency string) *Money {

	m := Money{}.create(value, currency)

	return m
}

func (m Money) ensureValidCurrency(currency string) error {
	if len(currency) != 3 {
		return ErrorMoneyCurrency
	}

	return nil
}

func (m Money) ensureValidAmount(amount float64) error {
	if amount <= 0 {
		return ErrorMoneyAmount
	}

	return nil
}

func (m Money) create(value float64, currency string) *Money {
	
	m.ensureValidCurrency(currency)
	m.ensureValidAmount(value)
	
	return &Money{
		amount: big.NewFloat(value),
		currency: currency,
	}
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Mul(amount float64) (*Money) {

    totalAmount := new(big.Float).Mul(m.amount, big.NewFloat(amount))
	
	m.amount = totalAmount
	return &m
}

func (m Money) AddMoney(amount *Money) (*Money, error) {
	
	if m.currency != amount.currency {
		return nil, ErrorMoneyCurrency
	}

    totalAmount := new(big.Float).Add(m.amount,  amount.amount)

	m.amount = totalAmount
	return &m, nil
}

func (m Money) Div(amount float64) (*Money) {
	
    totalAmount := new(big.Float).Quo(m.amount, big.NewFloat(amount))
	
	m.amount = totalAmount

	return &m 
}

func (m Money) Cents() int {
	cents, _ := new(big.Float).Mul(m.amount, big.NewFloat(100.0)).Int64()
	return int(cents)
}

func (m Money) Amount() float64 {
	v, _ := m.amount.Float64()
	return v
}

func (m *Money) String() string {
    if m == nil {
        return "<nil>"
    }
    return fmt.Sprintf("%.2f %s", m.Amount(), m.Currency())
}
