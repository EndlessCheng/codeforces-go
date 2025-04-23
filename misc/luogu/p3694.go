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
	pre := make([][20]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pre[i] = pre[i-1]
		pre[i][v-1]++
	}

	u := 1 << m
	subSum := make([]int, u)
	for i, v := range pre[n][:m] {
		highBit := 1 << i
		for mask, s := range subSum[:highBit] {
			subSum[highBit|mask] = s + v
		}
	}

	f := make([]int, u)
	for s, fs := range f {
		for cus, lb := u-1^s, 0; cus > 0; cus ^= lb {
			lb = cus & -cus
			ns := s | lb
			p := bits.TrailingZeros(uint(lb))
			f[ns] = max(f[ns], fs+pre[subSum[ns]][p]-pre[subSum[s]][p])
		}
	}
	Fprint(out, n-f[u-1])
}

//func main() { p3694(bufio.NewReader(os.Stdin), os.Stdout) }
