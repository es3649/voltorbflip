package main

import (
	// tm "github.com/buger/goterm"
	"fmt"

	"github.com/es3649/voltorbflip/internal/board"
	vb "github.com/es3649/voltorbflip/internal/voltorb"
)

func main() {
	vb.Voltorb()
	fmt.Println(12 / 8)
	// tm.Clear()
	// tm.MoveCursor(1, 1)
	b := board.Generate(1)
	b.Print()
}
