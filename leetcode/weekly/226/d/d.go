package main

import "math"

// github.com/EndlessCheng/codeforces-go
// 1278. 分割回文串 III
func palindromePartition(s string, k int) int {
	n := len(s)
	minChange := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		minChange[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			minChange[i][j] = minChange[i+1][j-1]
			if s[i] != s[j] {
				minChange[i][j]++
			}
		}
	}

	f := minChange[0]
	for i := 1; i < k; i++ {
		for r := n - k + i; r >= i; r-- {
			f[r] = math.MaxInt
			for l := i; l <= r; l++ {
				f[r] = min(f[r], f[l-1]+minChange[l][r])
			}
		}
	}
	return f[n-1]
}

func checkPartitioning(s string) bool {
	return palindromePartition(s, 3) == 0
}
