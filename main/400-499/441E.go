package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf441E(in io.Reader, out io.Writer) {
	var x, k int
	var p float64
	Fscan(in, &x, &k, &p)
	p /= 100
	f := make([]float64, k+1)
	for j := range f {
		f[j] = float64(bits.TrailingZeros(uint(x + j)))
	}
	for i := 1; i <= k; i++ {
		nf := make([]float64, k+1)
		nf[0] = (f[0] + 1) * p
		for j := 1; j <= k; j++ {
			nf[j-1] += f[j] * (1 - p)
			if j*2 <= k {
				nf[j*2] += (f[j] + 1) * p
			}
		}
		f = nf
	}
	Fprintf(out, "%.6f", f[0])
}

//func main() { cf441E(os.Stdin, os.Stdout) }
