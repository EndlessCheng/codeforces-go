package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p3694(in io.Reader, out io.Writer) {
	var n, m, v int
	Fscan(in, &n, &m)
	sum := make([][20]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		sum[i] = sum[i-1]
		sum[i][v-1]++
	}

	u := 1 << m
	sz := make([]int, u)
	for i, v := range sum[n][:m] {
		highBit := 1 << i
		for mask, s := range sz[:highBit] {
			sz[highBit|mask] = s + v
		}
	}

	f := make([]int, u)
	for s, fs := range f {
		for cus, lb := u-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			p := bits.TrailingZeros(uint(lb))
			f[ns] = max(f[ns], fs+sum[sz[ns]][p]-sum[sz[s]][p])
		}
	}
	Fprint(out, n-f[u-1])
}

//func main() { p3694(bufio.NewReader(os.Stdin), os.Stdout) }
