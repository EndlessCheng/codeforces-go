package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF463D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, ans int
	Fscan(in, &m, &n)
	a := make([]int, m)
	pos := make([][]int, n)
	for i := range pos {
		pos[i] = make([]int, m+1)
		for j := range a {
			Fscan(in, &a[j])
			pos[i][a[j]] = j
		}
	}

	dp := make([]int, m) // 最长路
	for i, v := range a {
	o:
		for j, w := range a[:i] {
			for _, p := range pos {
				if p[w] > p[v] { // 把其中一行当作（1-n），其余行映射一下
					continue o
				}
			}
			dp[i] = max(dp[i], dp[j])
		}
		dp[i]++
		ans = max(ans, dp[i])
	}
	Fprint(out, ans)
}

//func main() { CF463D(os.Stdin, os.Stdout) }
