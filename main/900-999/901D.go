package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf901D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, x, y, w, key int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	xs := make([]int, m+1)
	ys := make([]int, m+1)
	g := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &x, &y)
		xs[i], ys[i] = x, y
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n+1)
	dep := make([]int, n+1)
	fa := make([]int, n+1)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] {
				fa[y] = x
				dep[y] = dep[x] + 1
				dfs(y)
				a[x] -= a[y]
			}
		}
	}
	dfs(1)

	for i := 1; i <= m; i++ {
		x, y = xs[i], ys[i]
		if (dep[x]-dep[y])&1 == 0 {
			if dep[x] > dep[y] {
				x, y = y, x
			}
			w = a[1] / 2
			for j := y; j != x; j = fa[j] {
				if dep[j]&1 != 0 {
					a[j] += w
				} else {
					a[j] -= w
				}
			}
			for j := x; j != 0; j = fa[j] {
				if dep[j]&1 != 0 {
					a[j] += w * 2
				} else {
					a[j] -= w * 2
				}
			}
			if (dep[x]+1)&1 == 0 {
				w *= -1
			}
			key = i
			break
		}
	}

	if a[1] != 0 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := 1; i <= m; i++ {
		x, y = xs[i], ys[i]
		if dep[x] > dep[y] {
			x, y = y, x
		}
		if fa[y] == x {
			Fprintln(out, a[y])
		} else if i == key {
			Fprintln(out, w)
		} else {
			Fprintln(out, 0)
		}
	}
}

//func main() { cf901D(bufio.NewReader(os.Stdin), os.Stdout) }
