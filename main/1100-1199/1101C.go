package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1101C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type tuple struct{ l, r, i int }
		a := make([]tuple, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
			a[i].i = i
		}
		slices.SortFunc(a, func(a, b tuple) int { return a.l - b.l })

		maxR := a[0].r
		for i := 1; i < n; i++ {
			p := a[i]	
			if p.l > maxR {
				ans := make([]int, n)
				for j := range ans {
					ans[j] = 1
				}
				for _, p := range a[:i] {
					ans[p.i] = 2
				}
				for _, v := range ans {
					Fprint(out, v, " ")
				}
				Fprintln(out)
				continue o
			}
			maxR = max(maxR, p.r)
		}
		Fprintln(out, -1)
	}
}

//func main() { cf1101C(bufio.NewReader(os.Stdin), os.Stdout) }
