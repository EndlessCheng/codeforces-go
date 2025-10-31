package main

import (
	"cmp"
	. "fmt"
	"io"
	"math"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func p1429(in io.Reader, out io.Writer) {
	type vec struct{ x, y float64 }
	merge := func(a, b []vec) []vec {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]vec, 0, n+m)
		for {
			if i == n {
				return append(res, b[j:]...)
			}
			if j == m {
				return append(res, a[i:]...)
			}
			if a[i].y < b[j].y {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}
	var f func([]vec) float64
	f = func(ps []vec) float64 {
		n := len(ps)
		if n <= 1 {
			return math.MaxFloat64
		}
		m := n >> 1
		x := ps[m].x
		d := math.Min(f(ps[:m]), f(ps[m:]))
		copy(ps, merge(ps[:m], ps[m:]))
		checkPs := []vec{}
		for _, pi := range ps {
			if math.Abs(pi.x-x) > d+1e-8 {
				continue
			}
			for j := len(checkPs) - 1; j >= 0; j-- {
				pj := checkPs[j]
				dy := pi.y - pj.y
				if dy >= d {
					break
				}
				dx := pi.x - pj.x
				d = math.Min(d, math.Hypot(dx, dy))
			}
			checkPs = append(checkPs, pi)
		}
		return d
	}

	var n int
	Fscan(in, &n)
	ps := make([]vec, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	slices.SortFunc(ps, func(a, b vec) int { return cmp.Compare(a.x, b.x) })
	Fprintf(out, "%.4f", f(ps))
}

//func main() { p1429(bufio.NewReader(os.Stdin), os.Stdout) }
