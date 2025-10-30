package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf799C(in io.Reader, out io.Writer) {
	var n, c, d, x, p, ans int
	var s string
	Fscan(in, &n, &c, &d)
	type pair struct{ b, p int }
	g := [2][]pair{}
	for range n {
		Fscan(in, &x, &p, &s)
		b := s[0] - 'C'
		g[b] = append(g[b], pair{x, p})
	}

	a, b := g[0], g[1]
	maxB := []int{}
	f := func(a []pair, c int) {
		slices.SortFunc(a, func(a, b pair) int { return a.p - b.p })
		preMax := make([]int, len(a)+1)
		preMax[0] = -1e9
		mx := int(-1e9)
		j := 0
		for i, p := range a {
			d := c - p.p
			if d < 0 {
				break
			}
			preMax[i+1] = max(preMax[i], p.b)
			mx = preMax[i+1]
			if j < i && a[j].p <= d {
				j++
			}
			for j > 0 && a[j-1].p > d {
				j--
			}
			ans = max(ans, p.b+preMax[j])
		}
		maxB = append(maxB, mx)
	}
	f(a, c)
	f(b, d)
	Fprint(out, max(ans, maxB[0]+maxB[1]))
}

//func main() { cf799C(bufio.NewReader(os.Stdin), os.Stdout) }
