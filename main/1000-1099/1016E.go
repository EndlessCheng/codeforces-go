package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1016E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var sy, y int64
	var n, q, l, r, x int
	Fscan(in, &sy, &l, &r, &n)
	seg := make([]struct{ l, r int }, n)
	sum := make([]int64, n+1)
	for i := range seg {
		Fscan(in, &seg[i].l, &seg[i].r)
		sum[i+1] = sum[i] + int64(seg[i].r-seg[i].l)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x, &y)
		i := sort.Search(n, func(i int) bool { return int64(seg[i].r-l)*y > int64(seg[i].r-x)*sy })
		j := sort.Search(n, func(i int) bool { return int64(seg[i].l-r)*y >= int64(seg[i].l-x)*sy })
		if i == j {
			Fprintln(out, 0)
			continue
		}
		y1, y2 := float64(y), float64(y-sy)
		s := float64(sum[j]-sum[i]) / y1 * y2
		d := float64(x) - float64(x-seg[i].l)/y1*y2 - float64(l)
		if d < 0 {
			s += d
		}
		d = float64(x) + float64(seg[j-1].r-x)/y1*y2 - float64(r)
		if d > 0 {
			s -= d
		}
		Fprintf(out, "%.8f\n", s)
	}
}

//func main() { CF1016E(os.Stdin, os.Stdout) }
