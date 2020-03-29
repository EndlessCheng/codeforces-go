package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	var s []byte
	var k int8
	Fscan(bufio.NewReader(_r), &s, &k)
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(p int, sum int8, isUpper bool) int
	f = func(p int, sum int8, isUpper bool) (cnt int) {
		if sum == k {
			return 1
		}
		if p >= n {
			return 0
		}
		dv := &dp[p][sum]
		if !isUpper && *dv >= 0 {
			return *dv
		}
		defer func() {
			if !isUpper {
				*dv = cnt
			}
		}()
		up := byte('9')
		if isUpper {
			up = s[p]
		}
		cnt += f(p+1, sum, isUpper && '0' == up)
		for digit := byte('1'); digit <= up; digit++ {
			cnt += f(p+1, sum+1, isUpper && digit == up)
		}
		return
	}
	Fprint(_w, f(0, 0, true))
}

func main() { run(os.Stdin, os.Stdout) }
