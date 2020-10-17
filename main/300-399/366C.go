package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF366C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e4
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		if a[i] == k*b[i] {
			s += a[i]
		}
	}

	dp1 := make([]int, mx+1)
	for i := range dp1 {
		dp1[i] = -1e9
	}
	dp1[0] = 0
	for i, v := range a {
		if w := v - k*b[i]; w > 0 {
			for j := mx; j >= w; j-- {
				dp1[j] = max(dp1[j], dp1[j-w]+v)
			}
		}
	}
	dp2 := make([]int, mx+1)
	for i := range dp2 {
		dp2[i] = -1e9
	}
	dp2[0] = 0
	for i, v := range a {
		if w := k*b[i] - v; w > 0 {
			for j := mx; j >= w; j-- {
				dp2[j] = max(dp2[j], dp2[j-w]+v)
			}
		}
	}
	ans := 0
	for i, d := range dp1 {
		if d > 0 && dp2[i] > 0 {
			ans = max(ans, d+dp2[i])
		}
	}
	ans += s
	if ans == 0 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF366C(os.Stdin, os.Stdout) }
