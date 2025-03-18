package shared_ports

import domain "github.com/dacalin/simple_slot/engine/shared/domain"

type ISlotEngine interface {
	Spin(amount *domain.Money) (*domain.SpinResult, error)
}
