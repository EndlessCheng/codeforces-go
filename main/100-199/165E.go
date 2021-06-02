package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种用 lowbit 的思路见 https://codeforces.com/contest/165/submission/94728110

// github.com/EndlessCheng/codeforces-go
func CF165E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 22

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	dp := make([]int, 1<<mx)
	for i := range dp {
		dp[i] = -1
	}
	for i := range a {
		Fscan(in, &a[i])
		dp[a[i]] = a[i]
	}
	for i := 0; i < mx; i++ {
		for s := 0; s < 1<<mx; s++ {
			s |= 1 << i
			if dp[s^1<<i] > 0 {
				dp[s] = dp[s^1<<i]
			}
		}
	}
	for _, v := range a {
		Fprint(out, dp[1<<mx-1^v], " ")
	}
}

//func main() { CF165E(os.Stdin, os.Stdout) }
