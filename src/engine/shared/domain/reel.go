package domain

type Reel struct {
	id int
	symbols []string
	bias    []int
	totalBias int
}


func NewReel(id int) *Reel {
	return &Reel{id: id}
}

func (r Reel) Add(symbol string, bias int) *Reel {
	newSymbols := append(r.symbols, symbol)
	newBiases := append(r.bias, bias)

	var totalBias = 0

	for _, bias := range newBiases {
		totalBias += bias
	}

	return &Reel{symbols: newSymbols, bias: newBiases, totalBias: totalBias}
}

func (r Reel) TotalBias() int {
	return r.totalBias
}

func (r Reel) GetSymbolFromCumulativeBias(value int) string {
	var cumulativeBias = 0
	result := ""
	for i, bias := range r.bias {
		cumulativeBias += bias
		if cumulativeBias >= value {
			result = r.symbols[i]
			break
		}
	}

	return result
}
