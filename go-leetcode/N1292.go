package main

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	maxSide := min(m, n)

	prefixSum := make([][]int, m)
	for i := 0; i < m; i++ {
		prefixSum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				prefixSum[i][j] = mat[i][j]
				continue
			}
			if i == 0 {
				prefixSum[i][j] = prefixSum[i][j-1] + mat[i][j]
				continue
			}
			if j == 0 {
				prefixSum[i][j] = prefixSum[i-1][j] + mat[i][j]
				continue
			}
			prefixSum[i][j] = prefixSum[i][j-1] + prefixSum[i-1][j] - prefixSum[i-1][j-1] + mat[i][j]
		}
	}

	for side := maxSide; side > 0; side-- {
		for i := 0; i <= m-side; i++ {
			for j := 0; j <= n-side; j++ {
				val := prefixSum[i+side-1][j+side-1]
				if i > 0 {
					val -= prefixSum[i-1][j+side-1]
				}
				if j > 0 {
					val -= prefixSum[i+side-1][j-1]
				}
				if i > 0 && j > 0 {
					val += prefixSum[i-1][j-1]
				}
				if val <= threshold {
					return side
				}
			}
		}
	}
	return 0
}
