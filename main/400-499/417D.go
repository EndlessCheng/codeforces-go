package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf417D(in io.Reader, out io.Writer) {
	var n, m, b, t, v int
	Fscan(in, &n, &m, &b)
	type tuple struct{ x, k, m int }
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].k, &t)
		for range t {
			Fscan(in, &v)
			a[i].m |= 1 << (v - 1)
		}
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.k - b.k })

	ans := int(2e18)
	f := make([]int, 1<<m)
	for i := 1; i < 1<<m; i++ {
		f[i] = 2e18
	}
	for _, p := range a {
		for i := 1<<m - 1; i > 0; i-- {
			f[i] = min(f[i], f[i&^p.m]+p.x)
		}
		ans = min(ans, f[1<<m-1]+p.k*b)
	}
	if ans == 2e18 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf417D(bufio.NewReader(os.Stdin), os.Stdout) }
