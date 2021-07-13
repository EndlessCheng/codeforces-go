package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF75D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int64 = 1e9
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, m, k, id int
	var v int64
	Fscan(in, &n, &m)
	type res struct{ s, minS, maxS, maxSubS int64 }
	rs := make([]res, n)
	for i := range rs {
		r := res{0, inf, -inf, -inf}
		dp := -inf
		for Fscan(in, &k); k > 0; k-- {
			Fscan(in, &v)
			r.s += v
			r.minS = min(r.minS, r.s)
			r.maxS = max(r.maxS, r.s)
			dp = max(dp, 0) + v
			r.maxSubS = max(r.maxSubS, dp)
		}
		rs[i] = r
	}

	Fscan(in, &id)
	r := rs[id-1]
	ans := r.maxSubS
	minS := min(0, r.minS)
	base := r.s
	for m--; m > 0; m-- {
		Fscan(in, &id)
		r := rs[id-1]
		ans = max(ans, max(r.maxSubS, r.maxS+base-minS))
		minS = min(minS, base+r.minS)
		base += r.s
	}
	Fprint(out, ans)
}

//func main() { CF75D(os.Stdin, os.Stdout) }
