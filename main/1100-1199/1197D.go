package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1197D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n, m int
	var k, v, s, ans int64
	Fscan(in, &n, &m, &k)
	minS := make([]int64, m)
	for i := 1; i < m; i++ {
		minS[i] = 1e18
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		s += v*int64(m) - k
		for d := 0; d < m; d++ {
			ans = max(ans, s-minS[(i+d)%m]-k*int64(d))
		}
		minS[i%m] = min(minS[i%m], s)
	}
	Fprint(out, ans/int64(m))
}

//func main() { CF1197D(os.Stdin, os.Stdout) }
