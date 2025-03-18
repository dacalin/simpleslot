package domain

import (
	"fmt"
	"strings"
)

type SpinResult struct {
	VisibleReels *VisibleReels  // Visible symbols [reel][row] 
	WinningLines []*WinningLine  // Winning Lines [line_index][row] 
	WinAmount    *Money
	IsWin        bool
}

// String returns a formatted string representation of SpinResult
func (sr *SpinResult) String() string {
    var b strings.Builder
    
    // Display the visible reels
    b.WriteString("\n")
    if sr.VisibleReels != nil {
        b.WriteString(sr.VisibleReels.String())
    } else {
        b.WriteString("  <nil>\n")
    }
    
    // Show win status
    if sr.IsWin {
        b.WriteString("\n--------------------\n")
        b.WriteString(fmt.Sprintf("\nWIN! Amount: %s\n", sr.WinAmount.String()))
        b.WriteString("\n--------------------\n")

    } else {
        b.WriteString("\n--------------------\n")
        b.WriteString("\nNo win\n")
        b.WriteString("\n--------------------\n")
    }
    
    // Display winning lines if any
    if sr.IsWin && len(sr.WinningLines) > 0 {
        b.WriteString("\nWinning Lines:\n")
        for i, line := range sr.WinningLines {
            if line != nil {
                b.WriteString(fmt.Sprintf("  Line %d: %s\n", i, line))
            }
        }
    }
    
    return b.String()
}