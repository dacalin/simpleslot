package superwin

import (
	"log"
	domain "github.com/dacalin/simple_slot/engine/shared/domain"
	shared_ports "github.com/dacalin/simple_slot/shared/domain/ports"
)

const joker string = "X"

type SuperWinEngine struct {
	shared_ports.ISlotEngine
	reels        []*domain.Reel
	visibleRows  int
	reelsCount   int
	payTable     *domain.Paytable
	evalStrategy *EvalStrategy
	rng          shared_ports.RNG
}

func NewSuperWinEngine(version string, rng shared_ports.RNG) *SuperWinEngine {

	// Number of reels and rows. Could be done by configuration
	reelsCount := 4
	visibleRows := 3

	// Create 4 reels
	reels := make([]*domain.Reel, reelsCount)

	// Configure each reel
	for i := 0; i < reelsCount; i++ {
		reels[i] = domain.NewReel(i).Add("A", 120).Add("K", 100).Add("Q", 100).Add("J", 100).Add("10", 100).Add(joker, 140)
	}

	evalStrategy := NewEvalStrategy(payTable, joker)

	return &SuperWinEngine{
		reels:        reels,
		payTable:     payTable,
		visibleRows:  visibleRows,
		reelsCount:   reelsCount,
		evalStrategy: evalStrategy,
		rng:          rng,
	}

}

func (e *SuperWinEngine) Spin(amount *domain.Money) (*domain.SpinResult, error) {

	betAmount := amount.Div(float64(e.evalStrategy.GeLinesCount()))

	visibleReels := domain.NewVisibleReels()

	// We spin each reel to get the final display
	for i := 0; i < e.reelsCount; i++ {
		reel_result, err := e.spinReel(i)

		if err != nil {
			return nil, err
		}

		visibleReels.AddReelResult(reel_result)

	}

	// Now we have the final display, lets calculate the winning
	hasWinnings, winningLines := e.evalStrategy.EvaluateVisibleReel(visibleReels)

	// Calculate total win
	totalWinAmount := domain.NewMoney(0, amount.Currency())

	if hasWinnings {
		for i := range winningLines {
			multiplier := payTable.GetMultiplier(winningLines[i].WinSymbol)
			lineWinAmount := betAmount.Mul(float64(multiplier))
			totalWinAmount, _ = totalWinAmount.AddMoney(lineWinAmount)
		}
	}

	return &domain.SpinResult{
		VisibleReels: visibleReels,
		WinAmount:    totalWinAmount,
		WinningLines: winningLines,
		IsWin:        hasWinnings}, nil
}

func (e *SuperWinEngine) spinReel(reelIndex int) ([]string, error) {
	reel := e.reels[reelIndex]
	visibleRows := e.visibleRows
	result := make([]string, visibleRows)

	for row := 0; row < visibleRows; row++ {

		totalBias := reel.TotalBias()

		// Select base in bias
		randomValue, errRng := e.rng.Rand(0, int64(totalBias))
		if errRng != nil {
			log.Fatal(errRng.Error())
			return []string{}, nil
		}

		result[row] = reel.GetSymbolFromCumulativeBias(int(randomValue))
	}

	return result, nil
}
