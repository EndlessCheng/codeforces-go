package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1611E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, k, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		hasFriend := make([]bool, n)
		for ; k > 0; k-- {
			Fscan(in, &v)
			hasFriend[v-1] = true
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		minDep := make([]int, n)
		var build func(v, fa, d int)
		build = func(v, fa, d int) {
			minDep[v] = 1e9
			if hasFriend[v] {
				minDep[v] = d
			}
			for _, w := range g[v] {
				if w != fa {
					build(w, v, d+1)
					minDep[v] = min(minDep[v], minDep[w])
				}
			}
		}
		build(0, -1, 0)

		var f func(v, fa, d int) int
		f = func(v, fa, d int) (ans int) {
			if d*2 >= minDep[v] {
				return 1
			}
			for _, w := range g[v] {
				if w != fa {
					res := f(w, v, d+1)
					if res < 0 {
						return -1
					}
					ans += res
				}
			}
			if ans == 0 {
				return -1
			}
			return
		}
		Fprintln(out, f(0, -1, 0))
	}
}

//func main() { CF1611E2(os.Stdin, os.Stdout) }
