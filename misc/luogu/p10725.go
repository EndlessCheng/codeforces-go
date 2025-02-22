package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p10725(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) (int, int, int)
	dfs = func(v, fa int) (max1, max2, color1 int) {
		max2 = -1e18
		color1 = a[v]
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			m1, m2, c1 := dfs(w, v)
			m1++
			m2++
			if c1 != color1 {
				ans = max(ans, max1+m1)
				if m1 > max1 {
					max2 = max(max1, m2)
					max1, color1 = m1, c1
				} else {
					max2 = max(max2, m1)
				}
			} else {
				ans = max(ans, max1+m2, max2+m1)
				max1 = max(max1, m1)
				max2 = max(max2, m2)
			}
		}
		return
	}
	dfs(0, -1)
	Fprint(out, ans)
}

//func main() { p10725(bufio.NewReader(os.Stdin), os.Stdout) }
