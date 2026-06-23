package main

import "math"

// github.com/EndlessCheng/codeforces-go
func pathsWithMaxScore1(board []string) []int {
	const mod = 1_000_000_007
	m, n := len(board), len(board[0])
	maxSum := make([][]int, m+1) // 定义同 64 题（改成最大路径和）
	ways := make([][]int, m+1)   // 定义同 63 题
	for i := range maxSum {
		maxSum[i] = make([]int, n+1)
		for j := range maxSum[i] {
			maxSum[i][j] = math.MinInt
		}
		ways[i] = make([]int, n+1)
	}
	maxSum[0][0] = 0 // 为什么这样写？见 64 题我的题解
	ways[0][0] = 1

	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				continue
			}
			// 左上、正上、正左
			maxSum[i+1][j+1] = max(maxSum[i][j], maxSum[i][j+1], maxSum[i+1][j])
			s := maxSum[i+1][j+1]
			w := 0
			// 如果路径和相同，则累加方案数（加法原理）
			if maxSum[i][j] == s {
				w += ways[i][j]
			}
			if maxSum[i][j+1] == s {
				w += ways[i][j+1]
			}
			if maxSum[i+1][j] == s {
				w += ways[i+1][j]
			}
			ways[i+1][j+1] = w % mod
			if '1' <= ch && ch <= '9' {
				maxSum[i+1][j+1] += int(ch - '0') // 加上当前格子的值
			}
		}
	}

	if maxSum[m][n] < 0 {
		return []int{0, 0}
	}
	return []int{maxSum[m][n], ways[m][n]}
}

func pathsWithMaxScore(board []string) []int {
	const mod = 1_000_000_007
	n := len(board[0])
	maxSum := make([]int, n+1) // 定义同 64 题（改成最大路径和）
	for i := 1; i <= n; i++ {
		maxSum[i] = math.MinInt
	}
	ways := make([]int, n+1) // 定义同 63 题
	ways[0] = 1

	for _, row := range board {
		preS, preW := maxSum[0], ways[0]
		maxSum[0], ways[0] = math.MinInt, 0
		for j, ch := range row {
			if ch == 'X' {
				preS, preW = maxSum[j+1], ways[j+1]
				maxSum[j+1], ways[j+1] = math.MinInt, 0
				continue
			}
			tmpS, tmpW := maxSum[j+1], ways[j+1]
			// 左上、正上、正左
			s := max(preS, maxSum[j+1], maxSum[j])
			// 如果路径和相同，则累加方案数（加法原理）
			w := 0
			if preS == s {
				w += preW
			}
			if maxSum[j+1] == s {
				w += ways[j+1]
			}
			if maxSum[j] == s {
				w += ways[j]
			}
			ways[j+1] = w % mod
			maxSum[j+1] = s
			if '1' <= ch && ch <= '9' {
				maxSum[j+1] += int(ch - '0') // 加上当前格子的值
			}
			preS, preW = tmpS, tmpW
		}
	}

	if maxSum[n] < 0 {
		return []int{0, 0}
	}
	return []int{maxSum[n], ways[n]}
}
