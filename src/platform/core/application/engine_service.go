package application

import (
	"github.com/dacalin/simple_slot/engine/shared/domain"
	shared_ports "github.com/dacalin/simple_slot/shared/domain/ports"
)


type EngineService struct {
	rng shared_ports.RNG
	engine shared_ports.ISlotEngine
}

func NewEngineService(engine shared_ports.ISlotEngine, rng shared_ports.RNG) *EngineService {

	return &EngineService{engine: engine, rng: rng}
}

// PlaceBets
// Receive Bet Entity. Bets are idempotent. We can user a BetId UUIDv4 or UUIDv7
func (e *EngineService) PlaceBet(amount *domain.Money) (error) {
	
	// Place Bet
	// In operator Wallet
	// And Save the request in DB
	// 
	

	return nil
}

func (e *EngineService) GetBetResult(amount *domain.Money) (*domain.SpinResult, error) {
	
	ret, err := e.engine.Spin(amount)

	// Save ret in DB
	// repository.Save(ret)

	// Check Jackpots. 
	// Use rng for it

	return ret, err
}