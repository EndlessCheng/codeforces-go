package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1082C(in io.Reader, out io.Writer) {
	var n, m, s, r, ans, tot int
	Fscan(in, &n, &m)
	type pair struct {
		s int
		a []int
	}
	g := make([]pair, m+1)
	for range n {
		Fscan(in, &s, &r)
		g[s].a = append(g[s].a, r)
	}
	for _, p := range g {
		slices.SortFunc(p.a, func(a, b int) int { return b - a })
	}

	for i := 0; len(g) > 0; i++ {
		newG := []pair{}
		for _, p := range g {
			if len(p.a) <= i {
				tot -= p.s
				continue
			}
			tot += p.a[i]
			p.s += p.a[i]
			if p.s <= 0 {
				tot -= p.s
			} else {
				newG = append(newG, p)
			}
		}
		ans = max(ans, tot)
		g = newG
	}
	Fprint(out, ans)
}

//func main() { cf1082C(bufio.NewReader(os.Stdin), os.Stdout) }
