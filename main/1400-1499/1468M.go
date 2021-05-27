package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1468M(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		nn := n
		a := make([][]int, n)
		vid := map[int]int{}
		for i := range a {
			Fscan(in, &m)
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &v)
				if vid[v] == 0 {
					vid[v] = n
					n++
				}
				a[i][j] = vid[v]
			}
		}
		g := make([][]int, n)
		deg := make([]int, n)
		for v, r := range a {
			for _, w := range r {
				g[v] = append(g[v], w)
				g[w] = append(g[w], v)
				deg[v]++
				deg[w]++
			}
		}
		less := func(v, w int) bool { return deg[v] < deg[w] || deg[v] == deg[w] && v < w }

		g2 := make([][]int, n)
		for v, ws := range g {
			for _, w := range ws {
				if less(v, w) {
					g2[v] = append(g2[v], w)
				}
			}
		}
		id := make([]int, n)
		for v, ws := range g {
			for _, w := range ws {
				for _, u := range g2[w] {
					if less(v, u) {
						if id[u] > 0 {
							if v < nn {
								Fprintln(out, v+1, u+1)
							} else {
								Fprintln(out, w+1, id[u])
							}
							continue o
						}
						id[u] = w + 1
					}
				}
			}
			for _, w := range ws {
				for _, u := range g2[w] {
					if less(v, u) {
						id[u] = 0
					}
				}
			}
		}
		Fprintln(out, -1)
	}
}

// 分块解法
func CF1468M2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ v, i int }

	var T, n, m, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]int, n)
		vid := map[int]int{}
		for i := range a {
			Fscan(in, &m)
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &v)
				if _, has := vid[v]; !has {
					vid[v] = len(vid)
				}
				a[i][j] = vid[v]
			}
		}

		n = len(vid)
		sz := int(math.Round(math.Sqrt(float64(n)) / 3)) // 不能直接用 sqrt(n)，会 MLE，测试发现 /3 比较适合

		has := make([]int, n)
		for i, b := range a {
			if len(b) < sz {
				continue
			}
			for _, v := range b {
				has[v] = i + 1
			}
			for j, c := range a {
				if j == i {
					continue
				}
				found := false
				for _, v := range c {
					if has[v] == i+1 {
						if found {
							Fprintln(out, i+1, j+1)
							continue o
						}
						found = true
					}
				}
			}
		}

		// 这种写法比 map 更省内存
		groups := make([][]pair, n)
		for i, b := range a {
			if len(b) < sz {
				sort.Ints(b)
				for j, v := range b {
					for _, w := range b[:j] {
						groups[v] = append(groups[v], pair{w, i + 1})
					}
				}
			}
		}
		id := make([]int, n)
		for _, g := range groups {
			for _, p := range g {
				if id[p.v] > 0 {
					Fprintln(out, id[p.v], p.i)
					continue o
				}
				id[p.v] = p.i
			}
			for _, p := range g {
				id[p.v] = 0
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { CF1468M(os.Stdin, os.Stdout) }
