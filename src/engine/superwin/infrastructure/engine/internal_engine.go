package superwin_internal_adapter

import (
	"github.com/dacalin/simple_slot/engine/shared/domain"
	superwin "github.com/dacalin/simple_slot/engine/superwin/domain"
	shared_ports "github.com/dacalin/simple_slot/shared/domain/ports"
)

// This is just an internal implementation, 
// but we could implement a grpc or http adapter as well
type SuperWinInternalAdapter struct {
	shared_ports.ISlotEngine
	engine *superwin.SuperWinEngine
}

func NewSuperWinInternalAdapter(version string, rng shared_ports.RNG) *SuperWinInternalAdapter {
	engine := superwin.NewSuperWinEngine(version, rng)

	return &SuperWinInternalAdapter{engine: engine}
}

func (e *SuperWinInternalAdapter) Spin(amount *domain.Money) (*domain.SpinResult, error) {
	
	return e.engine.Spin(amount)
}