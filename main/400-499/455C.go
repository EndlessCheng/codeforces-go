package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf455C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, op, x, y int
	Fscan(in, &n, &m, &q)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	d := make([]int, n+1)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	for i := 1; i <= n; i++ {
		if fa[i] != i {
			continue
		}
		var dfs func(int, int) int
		dfs = func(x, from int) (maxL int) {
			fa[x] = i
			for _, y := range g[x] {
				if y != from {
					subL := dfs(y, x) + 1
					d[i] = max(d[i], maxL+subL)
					maxL = max(maxL, subL)
				}
			}
			return
		}
		dfs(i, 0)
	}

	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		if op == 1 {
			Fprintln(out, d[find(x)])
			continue
		}
		Fscan(in, &y)
		x = find(x)
		y = find(y)
		if x != y {
			d[x] = max(d[x], d[y], (d[x]+1)/2+(d[y]+1)/2+1)
			fa[y] = x
		}
	}
}

//func main() { cf455C(bufio.NewReader(os.Stdin), os.Stdout) }
