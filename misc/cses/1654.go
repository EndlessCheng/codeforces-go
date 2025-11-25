package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
	"slices"
)

// https://cses.fi/problemset/task/1654/
func p1654() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

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
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i
			f[s] += f[s^1<<i]
		}
	}

	g := make([]int, 1<<w)
	for _, v := range a {
		g[v]++
	}
	for i := range w {
		for s := 1<<w - 1; s >= 0; s-- {
			s &^= 1 << i
			g[s] += g[s|1<<i]
		}
	}

	for _, v := range a {
		Fprintln(out, f[v], g[v], n-f[1<<w-1^v])
	}
}
