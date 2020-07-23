package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF106C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxW, n, a, b int
	Fscan(in, &maxW, &n)
	n++
	weights := make([]int, n)
	values := make([]int, n)
	stocks := make([]int, n)
	Fscan(in, &weights[0], &values[0])
	stocks[0] = maxW / weights[0]
	for i := 1; i < n; i++ {
		Fscan(in, &a, &b, &weights[i], &values[i])
		stocks[i] = a / b
	}

	dp := make([]int, maxW+1)
	for i, v := range values {
		num, w := stocks[i], weights[i]
		for k := 1; num > 0; k <<= 1 {
			K := min(k, num)
			for j := maxW; j >= K*w; j-- {
				dp[j] = max(dp[j], dp[j-K*w]+K*v)
			}
			num -= K
		}
	}
	Fprint(out, dp[maxW])
}

//func main() { CF106C(os.Stdin, os.Stdout) }
