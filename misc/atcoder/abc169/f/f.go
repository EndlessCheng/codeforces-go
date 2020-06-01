package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, sum int
	Fscan(in, &n, &sum)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, v := range a {
		for s := sum; s >= 0; s-- {
			dp[s] <<= 1
			if s >= v {
				dp[s] += dp[s-v]
			}
			dp[s] %= 998244353
		}
	}
	Fprint(_w, dp[sum])
}

func main() { run(os.Stdin, os.Stdout) }
