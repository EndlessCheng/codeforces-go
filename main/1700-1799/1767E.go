package main

import (
	. "fmt"
	"io"
	"math/bits"
)

func cf1767E(in io.Reader, out io.Writer) {
	var n, m, v, w, s int
	Fscan(in, &n, &m, &v)
	v--
	g := make([]int, m)
	g[v] |= 1 << v
	for range n - 1 {
		Fscan(in, &w)
		w--
		g[v] |= 1 << w
		g[w] |= 1 << v
		v = w
	}
	g[v] |= 1 << v

	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}

	memo := map[int]int{0: 0}
	var f func(int) int
	f = func(s int) int {
		if v, ok := memo[s]; ok {
			return v
		}
		tz := bits.TrailingZeros(uint(s))
		lb := 1 << tz
		res := f(s ^ lb)
		if g[tz]&lb == 0 {
			res = max(res, f(s&^g[tz]^lb)+a[tz])
		}
		memo[s] = res
		return res
	}
	Fprint(out, s-f(1<<m-1))
}

//func main() { cf1767E(bufio.NewReader(os.Stdin), os.Stdout) }
