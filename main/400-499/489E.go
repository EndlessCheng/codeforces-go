package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF489E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var lim float64
	Fscan(in, &n, &lim)
	x := make([]float64, n+1)
	b := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &x[i], &b[i])
	}

	dp := make([]float64, n+1)
	from := make([]int, n+1)
	l, r := 0., 1e3
	for step := 40; step > 0; step-- {
		mid := (l + r) / 2
		for i := 1; i <= n; i++ {
			dp[i] = 1e99
			for j, dv := range dp[:i] {
				res := dv + math.Sqrt(math.Abs(x[i]-x[j]-lim)) - b[i]*mid
				if res < dp[i] {
					dp[i] = res
					from[i] = j
				}
			}
		}
		if dp[n] <= 0 {
			r = mid
		} else {
			l = mid
		}
	}

	ans := []int{}
	for i := n; i > 0; i = from[i] {
		ans = append(ans, i)
	}
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i], " ")
	}
}

//func main() { CF489E(os.Stdin, os.Stdout) }
