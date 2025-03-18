package shared_ports

type RNG interface {
	Rand(min int64, max int64) (value int64, err error)
}