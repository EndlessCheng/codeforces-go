package main

import (
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m int
	var s []byte
	Fscan(in, &n, &m, &s)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, 1<<n)
	}
	f[0][0] = 1
	for i := range m {
		for c := byte('a'); c <= 'z'; c++ {
			for diff, fv := range f[i] {
				var lcs, newLcs, newDiff int
				for j, b := range s {
					one := diff >> j & 1
					pre := newLcs
					newLcs = max(newLcs, lcs+(one|b2i(b == c)))
					newDiff |= (newLcs - pre) << j
					lcs += one
				}
				f[i+1][newDiff] = (f[i+1][newDiff] + fv) % mod
			}
		}
	}

	ans := make([]int, n+1)
	for diff, fv := range f[m] {
		ans[bits.OnesCount(uint(diff))] += fv
	}
	for _, v := range ans {
		Fprint(out, v%mod, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
