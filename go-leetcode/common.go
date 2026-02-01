package main

import "fmt"

func main() {
	fmt.Println(minimumPairRemoval([]int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
