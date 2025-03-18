package superwin

import (
	"fmt"
	"testing"
	"github.com/dacalin/simple_slot/engine/shared/domain"
	shared_rng "github.com/dacalin/simple_slot/shared/infrastructure/rng"
)


func TestEngineRTP(t *testing.T) {

	targetRTP := 0.977
	epsilon := 0.001

	amount := domain.NewMoney(1.0, "EUR")
	engine := NewSuperWinEngine("97", shared_rng.NewInternalRNG())

	numSpins := 10000000
	var totalBet float64 = 0
	var totalWon float64 = 0

	for  i := 0; i<numSpins; i++ {
		totalBet += amount.Amount()
		result, _ := engine.Spin(amount)
		totalWon += result.WinAmount.Amount()
	}

	RTP := totalWon/totalBet
	t.Log(fmt.Sprintf("TOTAL BET: %f", totalBet))
	t.Log(fmt.Sprintf("TOTAL WON: %f", totalWon))
	t.Log(fmt.Sprintf("TOTAL RTP: %f", RTP))

	if RTP > targetRTP + epsilon || RTP < targetRTP - epsilon {
		t.Errorf("TargetRTP: %f, RealRTP: %f", targetRTP, RTP)
	}	
}