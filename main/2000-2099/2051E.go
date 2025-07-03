package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2051E(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type pair struct{ v, d int }
		es := make([]pair, n*2)
		for i := range n {
			Fscan(in, &v)
			es[i] = pair{v, 1}
		}
		for i := range n {
			Fscan(in, &v)
			es[n+i] = pair{v, -1}
		}
		slices.SortFunc(es, func(a, b pair) int { return a.v - b.v })

		ans, customer, neg := 0, n, 0
		for i, p := range es {
			if (i == n*2-1 || p.v != es[i+1].v) && neg <= k {
				ans = max(ans, p.v*customer)
			}
			neg += p.d
			if p.d < 0 {
				customer--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2051E(bufio.NewReader(os.Stdin), os.Stdout) }
