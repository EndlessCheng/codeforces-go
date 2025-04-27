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
	minCost := make([]int, n)
	for i := range cost {
		cost[i] = make([]int, m)
		for j := range cost[i] {
			Fscan(in, &cost[i][j])
		}
		minCost[i] = slices.Min(cost[i])
	}

	type pair struct{ mask, cost int }
	same := make([][]pair, n)
	for i, s := range a {
		same[i] = make([]pair, m)
		for j, b := range s {
			mx := 0
			for k, t := range a {
				if t[j] == b {
					same[i][j].mask |= 1 << k
					same[i][j].cost += cost[k][j]
					mx = max(mx, cost[k][j])
				}
			}
			same[i][j].cost -= mx
		}
	}

	u := 1 << n
	f := make([]int, u)
	for i := 1; i < u; i++ {
		f[i] = 1e9
	}
	for i, fv := range f[:u-1] {
		j := bits.TrailingZeros(^uint(i))
		f[i|1<<j] = min(f[i|1<<j], fv+minCost[j]) // 单改 j
		for _, p := range same[j] {
			f[i|p.mask] = min(f[i|p.mask], fv+p.cost) // 改 mask 这一组
		}
	}
	Fprint(out, f[u-1])
}

//func main() { cf543C(os.Stdin, os.Stdout) }
