package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF835D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	dp := make([][]int, n+1)
	isP := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		isP[i] = make([]bool, n+1)
	}
	ans := make([]int, n+1)
	for sz := 1; sz <= n; sz++ {
		for l, r := 1, sz; r <= n; l++ {
			isP[l][r] = s[l-1] == s[r-1] && (r-l < 2 || isP[l+1][r-1])
			if isP[l][r] {
				dp[l][r] = dp[l][l+sz/2-1] + 1
				ans[dp[l][r]]++
			}
			r++
		}
	}
	for i := n; i > 1; i-- {
		ans[i-1] += ans[i]
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF835D(os.Stdin, os.Stdout) }
