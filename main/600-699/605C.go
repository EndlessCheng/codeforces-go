package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf605C(in io.Reader, out io.Writer) {
	var n, p, q int
	Fscan(in, &n, &p, &q)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		Fscan(in, &a[i], &b[i])
	}

	f := func(x float64) float64 {
		y := 1.
		for i, v := range a {
			y = min(y, (1-x*float64(v))/float64(b[i]))
		}
		return x*float64(p) + y*float64(q)
	}

	l, r := 0., 1/float64(slices.Max(a))
	for range 37 {
		m1 := l + (r-l)/3
		m2 := r - (r-l)/3
		if f(m1) > f(m2) {
			r = m2
		} else {
			l = m1
		}
	}
	Fprintf(out, "%.6f", f(l))
}

//func main() { cf605C(bufio.NewReader(os.Stdin), os.Stdout) }
