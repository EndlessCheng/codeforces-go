package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf385D(in io.Reader, out io.Writer) {
	var n, l, r, v int
	Fscan(in, &n, &l, &r)
	a := make([]struct{ x, y, cot float64 }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &v)
		a[i].x -= float64(l)
		a[i].cot = math.Tan(float64(90-v) / 180 * math.Pi)
	}

	f := make([]float64, 1<<n)
	for s, preLen := range f {
		for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			i := bits.TrailingZeros32(uint32(lb))
			t := a[i]
			x, y := t.x-preLen, t.y
			v := x + y*t.cot
			if v <= 0 {
				Fprint(out, r-l)
				return
			}
			f[ns] = max(f[ns], preLen+(x*x+y*y)/v)
		}
	}
	Fprintf(out, "%.6f", min(f[len(f)-1], float64(r-l)))
}

//func main() { cf385D(os.Stdin, os.Stdout) }
