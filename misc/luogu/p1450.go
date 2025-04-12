package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p1450(in io.Reader, out io.Writer) {
	var n, s int
	var w, a [4]int
	const mx int = 1e5
	f := [mx + 1]int{1}
	for i := range w {
		Fscan(in, &w[i])
		for j := w[i]; j <= mx; j++ {
			f[j] += f[j-w[i]]
		}
	}
	for Fscan(in, &n); n > 0; n-- {
		for i := range a {
			Fscan(in, &a[i])
			a[i]++
		}
		Fscan(in, &s)
		ans := 0
		for sub := 0; sub < 1<<len(a); sub++ {
			s := s
			for t := uint(sub); t > 0; t &= t - 1 {
				i := bits.TrailingZeros(t)
				s -= a[i] * w[i]
			}
			if s >= 0 {
				ans += (1 - bits.OnesCount(uint(sub))%2*2) * f[s]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { p1450(bufio.NewReader(os.Stdin), os.Stdout) }
