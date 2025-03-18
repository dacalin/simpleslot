package domain

type Paytable struct {
	multipliers map[string]int
}

func NewPayTable() *Paytable {
	return &Paytable{
		multipliers: make(map[string]int),
	}
}

func (pt *Paytable) GetMultiplier(symbol string) int {
	return pt.multipliers[symbol]
}

func (pt *Paytable) Add(symbol string, multiplier int) *Paytable {

	pt.multipliers[symbol] = multiplier

	return pt
}

// Get the Symbol with the biggest multiplier
func (pt *Paytable) GetBestSymbol() string {

	var max = 0
	var dominantSymbol string

	for symbol, value := range pt.multipliers {
		if value > max {
			max = value
			dominantSymbol = symbol
		}	
	}

	return dominantSymbol
}