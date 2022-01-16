package main

// github.com/EndlessCheng/codeforces-go
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i, q := range questions {
		f[i+1] = max(f[i+1], f[i])
		j := i + q[1] + 1
		if j > n {
			j = n
		}
		f[j] = max(f[j], f[i]+q[0])
	}
	return int64(f[n])
}

func max(a, b int) int { if b > a { return b }; return a }
