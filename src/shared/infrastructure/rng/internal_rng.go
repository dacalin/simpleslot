package shared_rng

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dacalin/simple_slot/engine/shared/domain"
	shared_ports "github.com/dacalin/simple_slot/shared/domain/ports"
)

type InternalRNG struct {
	shared_ports.RNG
	rng *rand.Rand 
}

func NewInternalRNG() *InternalRNG {
	seed := time.Now().UnixMilli()
	source := rand.NewSource(seed)
	fmt.Println("Seed:", seed)

	return &InternalRNG{
		rng: rand.New(source),
	}
}

func (r *InternalRNG) Rand(min int64, max int64) (int64, error) {
	if min >= max {
		return 0, domain.ErrorRng
	}

	value := r.rng.Int63n(max-min+1) + min
	return value, nil
}
