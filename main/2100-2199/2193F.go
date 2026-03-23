package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2193F(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, sx, sy, tx, ty int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &sx, &sy, &tx, &ty)
		type pair struct{ x, y int }
		a := make([]pair, n+1)
		for i := range n {
			Fscan(in, &a[i].x)
		}
		for i := range n {
			Fscan(in, &a[i].y)
		}
		a[n] = pair{tx, ty}
		slices.SortFunc(a, func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })

		fl, fr := 0, 0
		preL, preR := sy, sy
		for i := 0; i <= n; {
			l := a[i].y
			for i++; i <= n && a[i].x == a[i-1].x; i++ {
			}
			r := a[i-1].y
			fl, fr = min(fl+abs(preL-r), fr+abs(preR-r))+r-l, min(fl+abs(preL-l), fr+abs(preR-l))+r-l
			preL, preR = l, r
		}
		Fprintln(out, tx-sx+min(fl, fr))
	}
}

//func main() { cf2193F(bufio.NewReader(os.Stdin), os.Stdout) }
