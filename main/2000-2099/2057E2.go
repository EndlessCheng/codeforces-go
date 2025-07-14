package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2057E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, q, v, w, wt int
	var k int16
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &q)
		f := make([][][]int16, n)
		for i := range f {
			f[i] = make([][]int16, n)
			for j := range f[i] {
				f[i][j] = make([]int16, n)
				for k := range f[i][j] {
					if k != j {
						f[i][j][k] = 1e3
					}
				}
			}
		}

		type edge struct{ v, w, wt int }
		es := make([]edge, m)
		for i := range es {
			Fscan(in, &v, &w, &wt)
			v--
			w--
			es[i] = edge{v, w, wt}
			f[0][v][w] = 1
			f[0][w][v] = 1
		}
		slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

		for k := range n {
			for i := range n {
				for j := range n {
					f[0][i][j] = min(f[0][i][j], f[0][i][k]+f[0][k][j])
				}
			}
		}

		a := make([]int, n-1)
		p := 0
		for _, e := range es {
			v, w := e.v, e.w
			if f[p][v][w] == 0 {
				continue
			}
			for i := range f {
				for j := range f {
					f[p+1][i][j] = min(f[p][i][j], f[p][i][v]+f[p][w][j], f[p][i][w]+f[p][v][j])
				}
			}
			a[p] = e.wt
			p++
		}

		for range q {
			Fscan(in, &v, &w, &k)
			v--
			w--
			i := sort.Search(n-1, func(i int) bool { return f[i+1][v][w] < k })
			Fprint(out, a[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2057E2(bufio.NewReader(os.Stdin), os.Stdout) }
