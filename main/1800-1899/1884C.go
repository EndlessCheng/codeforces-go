package main

import (
	. "fmt"
	"io"
	"maps"
	"slices"
)

// https://github.com/EndlessCheng
func cf1884C(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
		}
		f := func(ban int) (res int) {
			d := map[int]int{}
			for _, p := range a {
				if p.l > ban || p.r < ban {
					d[p.l]++
					d[p.r+1]--
				}
			}
			s := 0
			for _, x := range slices.Sorted(maps.Keys(d)) {
				s += d[x]
				res = max(res, s)
			}
			return
		}
		Fprintln(out, max(f(1), f(m)))
	}
}

//func main() { cf1884C(bufio.NewReader(os.Stdin), os.Stdout) }
