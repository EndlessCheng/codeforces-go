package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf698C(in io.Reader, out io.Writer) {
	var n, k, cnt int
	Fscan(in, &n, &k)
	a := make([]float64, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > 0 {
			cnt++
		}
	}
	k = min(k, cnt)

	ans := make([]float64, n)
	f := make([]float64, 1<<n)
	f[0] = 1
	for s := range f {
		sum := 0.
		for t := uint(s); t > 0; t &= t - 1 {
			sum += a[bits.TrailingZeros(t)]
		}
		for i, v := range a {
			if s>>i&1 == 0 {
				f[s|1<<i] += f[s] * v / (1 - sum)
			}
		}
		if bits.OnesCount(uint(s)) == k {
			for t := uint(s); t > 0; t &= t - 1 {
				ans[bits.TrailingZeros(t)] += f[s]
			}
		}
	}

	for _, v := range ans {
		Fprintf(out, "%.6f ", v)
	}
}

//func main() { cf698C(os.Stdin, os.Stdout) }
