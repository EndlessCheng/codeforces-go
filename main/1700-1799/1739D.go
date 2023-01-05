package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1739D(_r io.Reader, _w io.Writer) {
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
		Fscan(in, &n, &k)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &v)
			g[v-1] = append(g[v-1], w)
		}
		Fprintln(out, sort.Search(n, func(mx int) bool {
			if mx == 0 {
				return false
			}
			cnt := 0
			var f func(int, bool) int
			f = func(v int, cut bool) (h int) {
				for _, w := range g[v] {
					h = max(h, f(w, v > 0))
				}
				h++
				if cut && h == mx {
					cnt++
					return 0
				}
				return
			}
			f(0, false)
			return cnt <= k
		}))
	}
}

//func main() { CF1739D(os.Stdin, os.Stdout) }
