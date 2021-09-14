package main

// github.com/EndlessCheng/codeforces-go
func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	sum := make([]int, n+1) // dp 前缀和
	sum[1] = 1
	for i := 1; i < n; i++ {
		sum[i+1] = sum[i]
		if i >= minJump && s[i] == '0' && sum[i-minJump+1] > sum[max(0, i-maxJump)] {
			sum[i+1]++
		}
	}
	return sum[n] > sum[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
