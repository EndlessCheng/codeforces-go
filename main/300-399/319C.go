package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF319C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	dp := make([]int64, n)
	slope := func(i, j int) float64 { return float64(dp[j]-dp[i]) / float64(b[i]-b[j]) }
	q := []int{0}
	for i := 1; i < n; i++ {
		for len(q) > 1 && slope(q[0], q[1]) < float64(a[i]) {
			q = q[1:]
		}
		dp[i] = dp[q[0]] + int64(b[q[0]])*int64(a[i])
		for len(q) > 1 && slope(q[len(q)-1], i) < slope(q[len(q)-2], q[len(q)-1]) {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	Fprint(out, dp[n-1])
}

//func main() { CF319C(os.Stdin, os.Stdout) }
