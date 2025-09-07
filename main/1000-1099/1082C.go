package main

import (
	. "fmt"
	"io"
	"maps"
	"slices"
)

// https://github.com/EndlessCheng
func cf1082C(in io.Reader, out io.Writer) {
	var n, m, s, r, ans, tot int
	Fscan(in, &n, &m)
	g := map[int][]int{}
	for range n {
		Fscan(in, &s, &r)
		g[s] = append(g[s], r)
	}
	for _, a := range g {
		slices.SortFunc(a, func(a, b int) int { return b - a })
	}

	sum := make([]int, m+1)
	for i := 0; len(g) > 0; i++ {
		for v, a := range g {
			if len(a) <= i {
				tot -= sum[v]
				delete(g, v)
				continue
			}
			tot += a[i]
			sum[v] += a[i]
			if sum[v] <= 0 {
				tot -= sum[v]
				delete(g, v)
			}
		}
		ans = max(ans, tot)
		g = maps.Clone(g)
	}
	Fprint(out, ans)
}

//func main() { cf1082C(bufio.NewReader(os.Stdin), os.Stdout) }
