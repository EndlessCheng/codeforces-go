package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, tot int
	Fscan(in, &n, &tot)
	a := make([]struct{ s, p float64; c int }, n)
	for i := range a {
		Fscan(in, &a[i].s, &a[i].c, &a[i].p)
		a[i].p /= 100
	}

	f := make([][1 << 8]float64, tot+1)
	for i := range f {
		for mask := range 1 << n {
			for s := uint8(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros8(s)
				if c := a[j].c; c <= i {
					f[i][mask] = max(f[i][mask], a[j].p*(f[i-c][mask^1<<j]+a[j].s)+(1-a[j].p)*f[i-c][mask])
				}
			}
		}
	}
	Fprintf(out, "%.6f", f[tot][1<<n-1])
}

func main() { run(os.Stdin, os.Stdout) }
