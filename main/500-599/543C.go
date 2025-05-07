package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf543C(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	cost := make([][]int, n)
	rowCost := make([]int, n)
	for i := range cost {
		cost[i] = make([]int, m)
		for j := range cost[i] {
			Fscan(in, &cost[i][j])
		}
		rowCost[i] = slices.Min(cost[i])
	}

	type pair struct{ t, cost int }
	colCost := make([][]pair, n)
	for i, s := range a {
		colCost[i] = make([]pair, m)
		for j, b := range s {
			mx := 0
			for k, t := range a {
				if t[j] == b {
					colCost[i][j].t |= 1 << k
					colCost[i][j].cost += cost[k][j]
					mx = max(mx, cost[k][j])
				}
			}
			colCost[i][j].cost -= mx
		}
	}

	u := 1 << n
	f := make([]int, u)
	for i := 1; i < u; i++ {
		f[i] = 1e9
	}
	for s, fs := range f[:u-1] {
		i := bits.TrailingZeros(^uint(s))
		f[s|1<<i] = min(f[s|1<<i], fs+rowCost[i])
		for _, p := range colCost[i] {
			f[s|p.t] = min(f[s|p.t], fs+p.cost)
		}
	}
	Fprint(out, f[u-1])
}

//func main() { cf543C(os.Stdin, os.Stdout) }
