package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf416E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i {
				g[i][j] = 1e9
			}
		}
	}
	for range m {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v][w] = wt
		g[w][v] = wt
	}
	wt := make([][]int, n)
	for i := range wt {
		wt[i] = slices.Clone(g[i])
	}
	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	for i, gi := range g {
		cnt := make([]int, n)
		for j := range n {
			if j == i || gi[j] == 1e9 {
				continue
			}
			for k := range n {
				if k != j && gi[k]+wt[k][j] == gi[j] {
					cnt[j]++
				}
			}
		}
		for j := i + 1; j < n; j++ {
			s := 0
			for k := range n {
				if gi[k]+g[k][j] == gi[j] {
					s += cnt[k]
				}
			}
			Fprint(out, s, " ")
		}
	}
}

//func main() { cf416E(bufio.NewReader(os.Stdin), os.Stdout) }
