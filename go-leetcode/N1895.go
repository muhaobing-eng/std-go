package main

func sum(elems ...int) int {
	var sum int
	for _, elem := range elems {
		sum += elem
	}
	return sum
}

func checkMagicSquare(grid [][]int, startX, startY, n int) bool {
	var baseLine int
	for i := 0; i < n; i++ {
		baseLine += grid[startX][startY+i]
	}

	var (
		upToDownSum int
		downToUpSum int
	)
	for i := 0; i < n; i++ {
		var (
			rowSum int
			colSum int
		)
		if i == 0 {
			rowSum = baseLine
		}
		for j := 0; j < n; j++ {
			if i != 0 {
				rowSum += grid[startX+i][startY+j]
			}
			colSum += grid[startX+j][startY+i]

			if i == j {
				upToDownSum += grid[startX+i][startY+j]
			}
			if i+j == n-1 {
				downToUpSum += grid[startX+i][startY+j]
			}
		}
		if rowSum != baseLine {
			return false
		}
		if colSum != baseLine {
			return false
		}
	}
	if upToDownSum != baseLine {
		return false
	}
	if downToUpSum != baseLine {
		return false
	}
	return true
}

func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	maxN := rows
	if cols < maxN {
		maxN = cols
	}
	if maxN <= 1 {
		return 1
	}

	for n := maxN; n > 1; n-- {
		for i := 0; i <= rows-n; i++ {
			for j := 0; j <= cols-n; j++ {
				if checkMagicSquare(grid, i, j, n) {
					return n
				}
			}
		}
	}
	return 1
}
