package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1613(in io.Reader, out io.Writer) {
	var n, m, v, w int
	Fscan(in, &n, &m)
	const mx = 63
	pa := make([][][mx]bool, n)
	g := make([][]bool, n)
	for i := range pa {
		pa[i] = make([][mx]bool, n)
		g[i] = make([]bool, n)
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		pa[v-1][w-1][0] = true
		g[v-1][w-1] = true
	}

	for i := 0; i < mx-1; i++ {
		for x, px := range pa {
			for t, xt := range px {
				if !xt[i] {
					continue
				}
				for y, ty := range pa[t] {
					if ty[i] {
						px[y][i+1] = true
						g[x][y] = true
					}
				}
			}
		}
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for step := 0; ; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == n-1 {
				Fprint(out, step)
				return
			}
			for w, ok := range g[v] {
				if ok && !vis[w] {
					vis[w] = true
					q = append(q, w)
				}
			}
		}
	}
}

//func main() { p1613(bufio.NewReader(os.Stdin), os.Stdout) }
