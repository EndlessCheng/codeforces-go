package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
	"slices"
)

// https://cses.fi/problemset/task/3141/
func p3141() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	const mod = 1_000_000_007

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	w := bits.Len(uint(slices.Max(a)))
	f := make([]int, 1<<w)
	for _, v := range a {
		f[v]++
	}
	for i := range w {
		for s := 1<<w - 1; s >= 0; s-- {
			s &^= 1 << i
			f[s] += f[s|1<<i]
		}
	}

	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
	for i, fv := range f {
		f[i] = pow2[fv] - 1
	}

	for i := range w {
		for s := 1<<w - 1; s >= 0; s-- {
			s &^= 1 << i
			f[s] = (f[s] - f[s|1<<i]) % mod
		}
	}

	for _, v := range f[:n+1] {
		Fprint(out, (v+mod)%mod, " ")
	}
}
