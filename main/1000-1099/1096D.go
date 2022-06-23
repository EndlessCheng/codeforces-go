package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1096D(_r io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	in := bufio.NewReader(_r)
	var v int64
	var s string
	Fscan(in, &v, &s)
	dp := [4]int64{}
	for _, c := range s {
		Fscan(in, &v)
		switch c {
		case 'h': dp[0] += v
		case 'a': dp[1] = min(dp[0], dp[1]+v)
		case 'r': dp[2] = min(dp[1], dp[2]+v)
		case 'd': dp[3] = min(dp[2], dp[3]+v)
		}
	}
	Fprint(out, dp[3])
}

//func main() { CF1096D(os.Stdin, os.Stdout) }
