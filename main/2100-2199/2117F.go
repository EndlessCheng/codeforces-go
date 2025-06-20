package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2117F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow2 := func(n int) int {
		res := 1
		for range n {
			res = res * 2 % mod
		}
		return res
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		g[0] = []int{-1}
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		cnt := 0
		for _, to := range g {
			if len(to) > 3 {
				cnt = 2
				break
			}
			if len(to) == 3 {
				cnt++
			}
		}
		if cnt > 1 {
			Fprintln(out, 0)
			continue
		}
		if cnt == 0 {
			Fprintln(out, pow2(n))
			continue
		}

		a := []int{}
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			size := 1
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				sz := dfs(w, v)
				size += sz
				if len(g[v]) == 3 {
					a = append(a, sz)
				}
			}
			return size
		}
		dfs(0, -1)
		x, y := a[0], a[1]
		if x == y {
			Fprintln(out, pow2(n-x*2+1))
		} else {
			Fprintln(out, pow2(n-min(x, y)*2-1)*3%mod)
		}
	}
}

//func main() { cf2117F(bufio.NewReader(os.Stdin), os.Stdout) }
