package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF476E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s, p []byte
	Fscan(bufio.NewReader(in), &s, &p)
	n, m := len(s), len(p)

	leftS := make([][26]int, n)
	for j := 0; j < 26; j++ {
		leftS[0][j] = -1
	}
	for i, b := range s {
		if i > 0 {
			leftS[i] = leftS[i-1]
		}
		leftS[i][b-'a'] = i
	}

	leftP := make([]int, n)
	for i := range leftP {
		leftP[i] = -1
	}
	for i := m - 1; i < n; i++ {
		if i > 0 {
			leftP[i] = leftP[i-1]
		}
		if s[i] != p[m-1] {
			continue
		}
		j, k := i, m-2
		for ; k >= 0 && j > 0; k-- {
			j = leftS[j-1][p[k]-'a']
		}
		if k < 0 {
			leftP[i] = j
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	for i, l := range leftP {
		for j := 0; j <= i+1; j++ {
			if j <= i {
				dp[i+1][j] = dp[i][j]
			}
			if j > 0 {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-1])
			}
			if l < 0 {
				continue
			}
			del := i - l + 1 - m
			if j >= del && j-del <= l {
				dp[i+1][j] = max(dp[i+1][j], dp[l][j-del]+1)
			}
		}
	}
	for _, v := range dp[n] {
		Fprint(out, v, " ")
	}
}

//func main() { CF476E(os.Stdin, os.Stdout) }
