package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF980E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v, w int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 20
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var build func(v, p, d int)
	build = func(v, p, d int) {
		pa[v][0] = p
		dep[v] = d
		for _, w := range g[v] {
			if w != p {
				build(w, v, d+1)
			}
		}
	}
	build(n-1, -1, 0)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}

	save := make([]bool, n)
	save[n-1] = true
	for i, left := n-2, n-1-k; i >= 0; i-- {
		if save[i] {
			continue
		}
		v := i
		for j := mx - 1; j >= 0; j-- {
			if p := pa[v][j]; p != -1 && !save[p] {
				v = p
			}
		}
		if d := dep[i] - dep[v] + 1; d <= left {
			left -= d
			for v := i; !save[v]; v = pa[v][0] {
				save[v] = true
			}
		}
	}
	for i, b := range save {
		if !b {
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { CF980E(os.Stdin, os.Stdout) }
