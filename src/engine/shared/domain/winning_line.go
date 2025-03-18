package domain

import (
	"fmt"
	"strings"
)

type WinningLine struct {
	LineSymbols []string  // Symbols in Line
	WinSymbol   string    // Símbolo ganador (incluyendo resuelto con comodines)
	LinePos   	[]int     // Line positions reel[row]
}

func (wl *WinningLine) String() string {
    if wl == nil {
        return "<nil>"
    }
    
    var b strings.Builder
    b.WriteString(fmt.Sprintf("Symbols: "))

    for i, symbol := range wl.LineSymbols {
        if i > 0 {
            b.WriteString(", ")
        }
        b.WriteString(fmt.Sprintf("%s", symbol))
    }

	return b.String()
}