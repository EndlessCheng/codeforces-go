package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1714G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, p, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type edge struct{ to, a, b int }
		g := make([][]edge, n)
		for w := 1; w < n; w++ {
			Fscan(in, &p, &a, &b)
			g[p-1] = append(g[p-1], edge{w, a, b})
		}

		ans := make([]int, n)
		s := []int64{0}
		var f func(int, int64)
		f = func(v int, sa int64) {
			ans[v] = sort.Search(len(s), func(i int) bool { return s[i] > sa }) - 1
			s = append(s, 0)
			for _, e := range g[v] {
				s[len(s)-1] = s[len(s)-2] + int64(e.b)
				f(e.to, sa+int64(e.a))
			}
			s = s[:len(s)-1]
		}
		f(0, 0)
		for _, v := range ans[1:] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1714G(os.Stdin, os.Stdout) }
