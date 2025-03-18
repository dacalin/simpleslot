package main

import (
	"fmt"

	"github.com/dacalin/simple_slot/bootstrap"
	"github.com/dacalin/simple_slot/engine/shared/domain"
)

func main() {

	// We create the core application service. To generalize this we need a different approach
	superWinEngineService := bootstrap.Start()

	var betAmount float64

	for {
		// Prompt for name
		fmt.Print("Enter bet amount (EUR): ")
		fmt.Scanln(&betAmount) // Reads user input

		// This should be a bet object, just a simplification
		amount := domain.NewMoney(betAmount, "EUR")
		fmt.Println(amount)

		// We call the application service.
		// First we place the bet
		superWinEngineService.PlaceBet(amount)
		// TODO: error handling after

		// Now the bet has been accepted, we call the engine to get the result
		result, _ := superWinEngineService.GetBetResult(amount)

		//TODO: Send Winnings 
		fmt.Println(result)
	}
}
