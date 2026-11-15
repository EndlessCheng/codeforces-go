package main

import (
	"cmp"
	. "fmt"
	"io"
	"math"
	"math/cmplx"
	"slices"
)

// https://github.com/EndlessCheng
func cf1841F(in io.Reader, out io.Writer) {
	var n, x, y, z, w int
	var u, v, s complex128
	Fscan(in, &n)
	type pair struct {
		x float64
		y complex128
	}
	q := make([]pair, 0, n*2)
	for range n {
		Fscan(in, &x, &z, &y, &w)
		x -= z
		y -= w
		u = complex(float64(x), float64(y))
		v = complex(float64(-x), float64(-y))
		if cmplx.Phase(u) > 0 {
			s += u
		}
		q = append(q, pair{cmplx.Phase(u), u}, pair{cmplx.Phase(v), v})
	}
	slices.SortFunc(q, func(a, b pair) int { return cmp.Compare(a.x, b.x) })

	ans := cmplx.Abs(s)
	ans *= ans
	for _, p := range q {
		s += p.y
		m := cmplx.Abs(s)
		ans = math.Max(ans, m*m)
	}
	Fprintf(out, "%.f", ans)
}

//func main() { cf1841F(bufio.NewReader(os.Stdin), os.Stdout) }
