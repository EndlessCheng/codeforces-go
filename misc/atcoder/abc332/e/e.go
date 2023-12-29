package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, d, total int
	Fscan(in, &n, &d)
	w := make([]int, n)
	for i := range w {
		Fscan(in, &w[i])
		total += w[i]
	}
	avg := float64(total) / float64(d)

	m := 1 << n
	d2 := make([]float64, m)
	for i := range d2 {
		s := 0
		for j := uint(i); j > 0; j &= j - 1 {
			s += w[bits.TrailingZeros(j)]
		}
		v := float64(s) - avg
		d2[i] = v * v
	}

	f := make([]float64, m)
	for i := 1; i < m; i++ {
		f[i] = 1e99
	}
	for i := 0; i < d; i++ {
		for s := m - 1; s >= 0; s-- {
			t := m - 1 ^ s
			for sub := t; sub > 0; sub = (sub - 1) & t {
				f[s|sub] = math.Min(f[s|sub], f[s]+d2[sub])
			}
			f[s] = 1e99
		}
	}
	Fprintf(out, "%.15f", f[m-1]/float64(d))
}

func main() { run(os.Stdin, os.Stdout) }
