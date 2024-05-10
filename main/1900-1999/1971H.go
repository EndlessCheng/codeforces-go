package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1971H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := [3][]int{}
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}

		g := make([][]int, n*2)
		addEdge := func(x, a, y, b int) {
			if x < 0 {
				x = -x
				a ^= 1
			}
			if y < 0 {
				y = -y
				b ^= 1
			}
			v, w := x-1+a*n, y-1+b*n
			g[v] = append(g[v], w)
		}
		for j := 0; j < n; j++ {
			for p := 0; p < 3; p++ {
				for q := 0; q < 3; q++ {
					if q != p {
						addEdge(a[p][j], 0, a[q][j], 1)
					}
				}
			}
		}

		sid := make([]int, len(g))
		dfn := make([]int, len(g))
		clock := 0
		st := []int{}
		inSt := make([]bool, len(g))
		var tarjan func(int) int
		tarjan = func(v int) int {
			clock++
			dfn[v] = clock
			lowV := clock
			st = append(st, v)
			inSt[v] = true
			for _, w := range g[v] {
				if dfn[w] == 0 {
					lowW := tarjan(w)
					lowV = min(lowV, lowW)
				} else if inSt[w] {
					lowV = min(lowV, dfn[w])
				}
			}
			if dfn[v] == lowV {
				for {
					w := st[len(st)-1]
					st = st[:len(st)-1]
					inSt[w] = false
					sid[w] = v
					if w == v {
						break
					}
				}
			}
			return lowV
		}
		for i, ts := range dfn {
			if ts == 0 {
				tarjan(i)
			}
		}
		for i, id := range sid[:n] {
			if id == sid[i+n] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1971H(bufio.NewReader(os.Stdin), os.Stdout) }
