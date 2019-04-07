// algorithm.go contians the algorithm for filling the board

package board

// score minimum shall be 27*2^(level-1)
// maximum (not inclusive!) shall be 54*2^(level-1)

func genScoreTiles(level int) []int {
	// this will be the sum of the multiples we use
	sum := 2 * (level + 4)
	var scoreVals []int

	// the available score tiles shall be
	for i := 2; i <= 2+(level/8); i++ {
		scoreVals = append(scoreVals, i)
	}

	var tileList []int

	for sum >= (level / 8) {
		tile := scoreVals[0] // TODO some randome element from score vals

		sum -= tile
		tileList = append(tileList, tile)
	}

	return tileList
}

func genVoltorbs(level, boardSize int, tileList []int) []int {
	maxVolts := level + 5 // TODO plus randint(0,2)
	if boardSize-len(tileList) < maxVolts {
		maxVolts = boardSize - len(tileList)
	}

	var volts []int
	for len(volts) <= maxVolts {
		volts = append(volts, 0)
	}

	return volts
}
