package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf30C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type tuple struct {
		x, y, t int
		p       float64
	}
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].t, &a[i].p)
	}
	slices.SortFunc(a, func(a, b tuple) int { return a.t - b.t })

	f := make([]float64, n)
	for i, t := range a {
		for j, q := range a[:i] {
			if (q.x-t.x)*(q.x-t.x)+(q.y-t.y)*(q.y-t.y) <= (q.t-t.t)*(q.t-t.t) {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += t.p
	}
	Fprintf(out, "%.10f", slices.Max(f))
}

//func main() { cf30C(bufio.NewReader(os.Stdin), os.Stdout) }
