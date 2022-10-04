package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF16E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	m := 1 << n
	f := make([]float64, m)
	f[m-1] = 1
	for s := uint(m) - 1; s > 0; s-- {
		k := bits.OnesCount(s)
		if k == 1 {
			continue
		}
		c := 1 / float64(k*(k-1)/2)
		for x := s; x > 0; x &= x - 1 {
			i := bits.TrailingZeros(x)
			t := s ^ 1<<i
			p := 0.
			for y := t; y > 0; y &= y - 1 {
				p += a[bits.TrailingZeros(y)][i]
			}
			f[t] += f[s] * p * c
		}
	}
	for i := 1; i < m; i <<= 1 {
		Fprintf(out, "%.6f ", f[i])
	}
}

//func main() { CF16E(os.Stdin, os.Stdout) }
