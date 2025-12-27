package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1842F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n+1)
	for i := range n {
		type pair struct{ v, fa int }
		q := []pair{{i, -1}}
		s, j := 0, 0
		for d := 0; q != nil; d++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				s += d
				j++
				ans[j] = max(ans[j], (n-1)*j-s*2)
				for _, w := range g[p.v] {
					if w != p.fa {
						q = append(q, pair{w, p.v})
					}
				}
			}
		}
	}

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf1842F(bufio.NewReader(os.Stdin), os.Stdout) }
