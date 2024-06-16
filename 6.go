package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func minTripsDP(boxes []int, machines int, capacity int) int {
	m := len(boxes)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt // Initialize with a large value
		}
	}

	// Base cases
	for j := 0; j <= capacity; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= capacity; j++ {
			if boxes[i-1] <= j {
				dp[i][j] = min(dp[i-1][j], 1+dp[i-1][j-boxes[i-1]]) // Include or skip box
			} else {
				dp[i][j] = dp[i-1][j] // Skip box (doesn't fit)
			}
		}
	}

	return dp[m][capacity]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, k, m int
		fmt.Fscan(reader, &n, &k)

		fmt.Fscan(reader, &m)
		boxes := make([]int, m)
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscan(reader, &a)
			boxes[j] = 1 << a // Calculate box weight as 2^a
		}

		trips := minTripsDP(boxes, n, k)
		fmt.Fprintln(writer, trips)
	}
}
