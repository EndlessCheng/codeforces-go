package main

// github.com/EndlessCheng/codeforces-go
func canMakePaliQueries(s string, queries [][]int) (ans []bool) {
	n := len(s)
	sum := make([][26]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i]
		sum[i+1][b-'a']++
	}
	ans = make([]bool, len(queries))
	for i, q := range queries {
		l, r, c := q[0], q[1]+1, 0
		for j := 0; j < 26; j++ {
			c += (sum[r][j] - sum[l][j]) & 1
		}
		ans[i] = c-1 <= 2*q[2]
	}
	return
}
