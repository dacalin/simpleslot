package superwin

import "github.com/dacalin/simple_slot/engine/shared/domain"

var payTable = domain.NewPayTable()

func init() {
	payTable.Add("A", 20)
	payTable.Add("K", 15)
	payTable.Add("Q", 10)
	payTable.Add("J", 5)
	payTable.Add("10", 2)
	payTable.Add(joker, 0)
}
