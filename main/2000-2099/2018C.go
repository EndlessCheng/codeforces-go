package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2018C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		deg := make([]int, n)
		for i, to := range g {
			deg[i] = len(to)
		}
		deg[0] = 1e9

		maxSave := 0
		cnt := 0
		pa := make([]int, n)
		for i := 1; i < n; i++ {
			pa[i] = -1
		}
		q := []int{0}
		for len(q) > 0 {
			cnt += len(q)
			maxSave = max(maxSave, cnt)
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, w := range g[v] {
					if pa[w] < 0 {
						pa[w] = v
						q = append(q, w)
					}
				}
				for deg[v] == 1 {
					cnt--
					v = pa[v]
				}
				deg[v]--
			}
		}
		Fprintln(out, n-maxSave)
	}
}

//func main() { cf2018C(bufio.NewReader(os.Stdin), os.Stdout) }
