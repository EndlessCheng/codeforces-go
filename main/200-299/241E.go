package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf241E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	x := make([]int, m)
	y := make([]int, m)
	G := make([][]int, n+1)
	RG := make([][]int, n+1)
	for i := range m {
		Fscan(in, &x[i], &y[i])
		G[x[i]] = append(G[x[i]], y[i])
		RG[y[i]] = append(RG[y[i]], x[i])
	}

	vis := make([][2]bool, n+1)
	var dfs func(int, int)
	dfs = func(v, tp int) {
		vis[v][tp] = true
		var g []int
		if tp == 0 {
			g = G[v]
		} else {
			g = RG[v]
		}
		for _, w := range g {
			if !vis[w][tp] {
				dfs(w, tp)
			}
		}
	}
	dfs(1, 0)
	dfs(n, 1)

	dis := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dis[i] = 1e9
	}

	loop := 0
	for {
		loop++
		if loop > n {
			Fprint(out, "No")
			return
		}
		changed := false
		for i := range m {
			X, Y := x[i], y[i]
			if vis[X][0] && vis[Y][1] {
				if dis[Y] > dis[X]+2 {
					dis[Y] = dis[X] + 2
					changed = true
				}
				if dis[X] > dis[Y]-1 {
					dis[X] = dis[Y] - 1
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	Fprintln(out, "Yes")
	for i := range m {
		X, Y := x[i], y[i]
		if vis[X][0] && vis[Y][1] {
			Fprintln(out, dis[Y]-dis[X])
		} else {
			Fprintln(out, 1)
		}
	}
}

//func main() { cf241E(bufio.NewReader(os.Stdin), os.Stdout) }
