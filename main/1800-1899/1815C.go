package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1815C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			if v > 0 {
				g[w] = append(g[w], v)
			}
		}
		vis := make([]bool, n)
		vis[0] = true
		level := [][]int{}
		q := []int{0}
		left := n - 1
		for i := 0; len(q) > 0; i++ {
			level = append(level, append([]int{}, q...))
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, w := range g[v] {
					if !vis[w] {
						vis[w] = true
						q = append(q, w)
						left--
					}
				}
			}
		}
		if left > 0 {
			Fprintln(out, "INFINITE")
			continue
		}
		Fprintln(out, "FINITE")
		ans := []int{}
		k := len(level)
		for sz := k; sz > 0; sz-- {
			for i := k - 1; i >= k-sz; i-- {
				ans = append(ans, level[i]...)
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1815C(os.Stdin, os.Stdout) }
