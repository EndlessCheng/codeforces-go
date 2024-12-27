package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf744A(in io.Reader, out io.Writer) {
	var n, m, k, cv, ce, maxV, ans int
	Fscan(in, &n, &m, &k)
	a := make([]int, k)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
	}
	g := make([][]int, n)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		cv++
		ce += len(g[v])
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}

	leftV, leftE := n, m
	for _, v := range a {
		cv, ce = 0, 0
		dfs(v)
		leftV -= cv
		leftE -= ce / 2
		maxV = max(maxV, cv)
		ans += cv*(cv-1)/2 - ce/2
	}
	Fprint(out, ans+leftV*(leftV-1)/2-leftE+leftV*maxV)
}

//func main() { cf744A(bufio.NewReader(os.Stdin), os.Stdout) }
