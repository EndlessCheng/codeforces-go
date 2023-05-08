package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1801C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := range g {
			for Fscan(in, &k); k > 0; k-- {
				Fscan(in, &v)
				if g[i] == nil || v > g[i][len(g[i])-1] {
					g[i] = append(g[i], v)
				}
			}
		}
		sort.Slice(g, func(i, j int) bool { return g[i][len(g[i])-1] < g[j][len(g[j])-1] })
		f := make([]int, n+1)
		for i, a := range g {
			f[i+1] = f[i]
			for j, v := range a {
				k := sort.Search(i, func(i int) bool { return g[i][len(g[i])-1] >= v })
				f[i+1] = max(f[i+1], f[k]+len(a)-j)
			}
		}
		Fprintln(out, f[n])
	}
}

//func main() { CF1801C(os.Stdin, os.Stdout) }
