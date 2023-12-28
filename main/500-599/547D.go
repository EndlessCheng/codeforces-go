package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf547D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	g := make([][]int, n+1)
	var preX, preY [2e5 + 1]int
	for i := 1; i <= n; i++ {
		Fscan(in, &x, &y)
		if preX[x] == 0 {
			preX[x] = i
		} else {
			g[i] = append(g[i], preX[x])
			g[preX[x]] = append(g[preX[x]], i)
			preX[x] = 0
		}
		if preY[y] == 0 {
			preY[y] = i
		} else {
			g[i] = append(g[i], preY[y])
			g[preY[y]] = append(g[preY[y]], i)
			preY[y] = 0
		}
	}

	colors := make([]int, n+1)
	var dfs func(int, int)
	dfs = func(v, c int) {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == 0 {
				dfs(w, 3^c)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if colors[i] == 0 {
			dfs(i, 1)
		}
	}

	for _, c := range colors[1:] {
		Fprintf(out, "%c", " br"[c])
	}
}

//func main() { cf547D(os.Stdin, os.Stdout) }
