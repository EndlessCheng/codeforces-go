package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF629C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1e9 + 7

	var n, m, sum, min int
	var s []byte
	Fscan(in, &n, &m, &s)
	for _, b := range s {
		if b == '(' {
			sum++
		} else {
			if sum--; sum < min {
				min = sum
			}
		}
	}
	d := n - m
	if d+min < 0 {
		Fprint(out, 0)
		return
	}

	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, d+1)
	}
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 0; j <= d; j++ {
			if j > 0 {
				dp[i][j] = dp[i-1][j-1]
			}
			if j < d {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}

	ans := int64(0)
	for i, dv := range dp {
		for j := -min; j <= d && j+sum <= d; j++ {
			ans = (ans + int64(dv[j])*int64(dp[d-i][j+sum])) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { CF629C(os.Stdin, os.Stdout) }
