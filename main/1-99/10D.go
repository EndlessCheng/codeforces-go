package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF10D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ansJ int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &m)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	dp := make([][]int, n+1)
	fa := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
		fa[i] = make([]int, m)
	}
	for i, v := range a {
		mx, k := 0, -1
		for j, w := range b {
			if v == w {
				dp[i+1][j] = mx + 1
				fa[i+1][j] = k
			} else {
				dp[i+1][j] = dp[i][j]
				fa[i+1][j] = j
			}
			if w < v && dp[i][j] > mx {
				mx, k = dp[i][j], j
			}
		}
	}
	for j, dv := range dp[n] {
		if dv > dp[n][ansJ] {
			ansJ = j
		}
	}
	Fprintln(out, dp[n][ansJ])
	var print func(i, j int)
	print = func(i, j int) {
		if i == 0 || j < 0 {
			return
		}
		print(i-1, fa[i][j])
		if fa[i][j] < j {
			Fprint(out, b[j], " ")
		}
	}
	print(n, ansJ)
}

//func main() { CF10D(os.Stdin, os.Stdout) }
