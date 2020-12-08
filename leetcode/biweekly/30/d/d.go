package main

// github.com/EndlessCheng/codeforces-go
var win [1e5 + 1]bool

func init() {
	for i := 1; i <= 1e5; i++ {
		for j := 1; j*j <= i; j++ {
			if !win[i-j*j] {
				win[i] = true
				break
			}
		}
	}
}

func winnerSquareGame(n int) bool {
	return win[n]
}
