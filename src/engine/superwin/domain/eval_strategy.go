package superwin

import (
	"log"
	"github.com/dacalin/simple_slot/engine/shared/domain"
)


type EvalStrategy struct {
	jokerSymbol string
	payTable    *domain.Paytable
	winLines    [][]int    // [line][reel] = row
}


func NewEvalStrategy(payTable *domain.Paytable, jokerSymbol string) *EvalStrategy {

	defaultLines := [][]int {
		{1, 1, 1, 1}, 
	}
	
	return &EvalStrategy{
		payTable:    payTable,
		winLines:    defaultLines,
		jokerSymbol: jokerSymbol,
	}
}

func (s *EvalStrategy) GeLinesCount()  int {
	return len(s.winLines)
}

func (s *EvalStrategy) AddCustomLine(line []int) {
	s.winLines = append(s.winLines, line)
}

// Evaluate all lines and return results
func (s *EvalStrategy) EvaluateVisibleReel(reels *domain.VisibleReels) (hasWinnings bool, lines []*domain.WinningLine) {
	results := []*domain.WinningLine{}
	hasWinnings = false

	// Evaluate line by line
	for lineIdx, lineDefinition := range s.winLines {

		lineSymbols := []string{}
	
		var err error
		var symbol string

		for reelId, rowId := range lineDefinition {
			symbol, err = reels.Get(reelId, rowId)
			if err != nil {
				log.Fatal(err)
			}
			lineSymbols = append(lineSymbols, symbol)
		
		}
		
		//  Check if is a winning Line
		isWin, winSymbol := s.isWinningLine(lineSymbols)
		
		if isWin == true {
			hasWinnings = true
			results = append(results, &domain.WinningLine{
				LineSymbols: 	lineSymbols,
				WinSymbol:   	winSymbol,
				LinePos:   		s.winLines[lineIdx],
			})
		}


	}
	
	return hasWinnings, results
}

// Evaluate if the line has winning
func (s *EvalStrategy) isWinningLine(symbols []string) (bool, string) {

	lineLen := len(symbols)

	// First we count symbols frequency
	symbolCount := make(map[string]int)
	jokerCount := 0
	dominantSymbol := ""

	for _, symbol := range symbols {

		if symbol == s.jokerSymbol {
			jokerCount++
		} else {
			symbolCount[symbol]++
			dominantSymbol = symbol
		}
	}

	// A line formed by Jokers
	if jokerCount == lineLen {
		return true, s.payTable.GetBestSymbol()
	}

	// A line formed by symbols and Jokers
	if symbolCount[dominantSymbol] + jokerCount == lineLen {
		return true, dominantSymbol
	}

	// There was no line
	return false, dominantSymbol
}