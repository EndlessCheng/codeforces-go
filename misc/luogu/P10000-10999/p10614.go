package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p10614(in io.Reader, out io.Writer) {
	b2i := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	const mod = 1_000_000_007
	t := []byte("ACGT")
	var T, m int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &m)
		n := len(s)
		nxt := make([][4]int, 1<<n)
		for diff := range nxt {
			for i, c := range t {
				var lcs, newLcs, newDiff int
				for j, b := range s {
					one := diff >> j & 1
					pre := newLcs
					newLcs = max(newLcs, lcs+(one|b2i(b == c)))
					newDiff |= (newLcs - pre) << j
					lcs += one
				}
				nxt[diff][i] = newDiff
			}
		}

		f := make([]int, 1<<n)
		f[0] = 1
		for i := 0; i < m; i++ {
			nf := make([]int, 1<<n)
			for diff, fv := range f {
				for _, newDiff := range nxt[diff] {
					nf[newDiff] = (nf[newDiff] + fv) % mod
				}
			}
			f = nf
		}

		ans := make([]int, n+1)
		for diff, fv := range f {
			ans[bits.OnesCount(uint(diff))] += fv
		}
		for _, v := range ans {
			Fprintln(out, v%mod)
		}
	}
}

//func main() { p10614(bufio.NewReader(os.Stdin), os.Stdout) }
