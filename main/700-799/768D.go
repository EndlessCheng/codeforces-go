package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF768D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var k, q, p int
	Fscan(in, &k, &q)
	dp := make([]float64, k+1)
	dp[0] = 1
	ans := [1001]int{}
	for p, day := 1, 1; p < 1001; day++ {
		for i := k; i > 0; i-- {
			dp[i] = (dp[i]*float64(i) + dp[i-1]*float64(k-i+1)) / float64(k)
		}
		for ; p < 1001 && dp[k]*2000 > float64(p)-1e-7; p++ {
			ans[p] = day
		}
		dp[0] = 0
	}
	for ; q > 0; q-- {
		Fscan(in, &p)
		Fprintln(out, ans[p])
	}
}

//func main() { CF768D(os.Stdin, os.Stdout) }
