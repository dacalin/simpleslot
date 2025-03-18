package domain

import "strings"

type VisibleReels struct {
	visibleReels [][]string
	reelsCount  int
}

func NewVisibleReels() *VisibleReels {
	return &VisibleReels{}
}

func (r *VisibleReels) AddReelResult(result []string) *VisibleReels {
	r.visibleReels = append(r.visibleReels, result)
	r.reelsCount += 1

	return r
}

func (r *VisibleReels) ReelsCount() int {
	return r.reelsCount
}

func (r *VisibleReels) Get(reelId int, rowId int) (string, error) {

	if r.reelsCount < 1 {
		return "", ErrorVisibleReel
	}

	if reelId >= r.ReelsCount() || rowId >= len(r.visibleReels[0]) {
		return "", ErrorVisibleReel
	}

	return r.visibleReels[reelId][rowId], nil
}

func (vr *VisibleReels) String() string {
    if vr == nil || vr.reelsCount == 0 {
        return "  <nil>"
    }
    
    var b strings.Builder
    
    // Get the number of rows from the first reel
    rowCount := len(vr.visibleReels[0])
    
    // Find the maximum width needed for any symbol to align columns
    maxWidth := 0
    for reel := 0; reel < vr.reelsCount; reel++ {
        for row := 0; row < rowCount; row++ {
            symbol, err := vr.Get(reel, row)
            if err == nil && len(symbol) > maxWidth {
                maxWidth = len(symbol)
            }
        }
    }
    
    // Add some padding
    maxWidth += 1
    
    // Horizontal border line
    horizontalLine := "  +" + strings.Repeat("-", maxWidth*vr.reelsCount+(vr.reelsCount-1)) + "+\n"
    
    // Print top border
    b.WriteString(horizontalLine)
    
    // Print rows
    for row := 0; row < rowCount; row++ {
        b.WriteString("  |")
        for reel := 0; reel < vr.reelsCount; reel++ {
            symbol, err := vr.Get(reel, row)
            if err != nil {
                symbol = "?"
            }
            
            // Padding for alignment
            padding := maxWidth - len(symbol)
            b.WriteString(symbol)
            b.WriteString(strings.Repeat(" ", padding))
        }
        b.WriteString("|\n")
    }
    
    // Print bottom border
    b.WriteString(horizontalLine)
    
    return b.String()
}