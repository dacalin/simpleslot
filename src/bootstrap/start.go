package bootstrap

import (
	superwin_internal_adapter "github.com/dacalin/simple_slot/engine/superwin/infrastructure/engine"
	"github.com/dacalin/simple_slot/platform/core/application"
	shared_rng "github.com/dacalin/simple_slot/shared/infrastructure/rng"
)

func Start() *application.EngineService {
	// Create Engine Adapter. In this case is just a wrapper for the domain object.
	engine_adapter := superwin_internal_adapter.NewSuperWinInternalAdapter("97", shared_rng.NewInternalRNG())

	// We inject into the core application service
	return application.NewEngineService(engine_adapter, shared_rng.NewInternalRNG())
}
