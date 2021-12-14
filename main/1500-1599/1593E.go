package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1593E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]int, n)
		d := make([]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
			d[v]++
			d[w]++
		}
		if n < 3 {
			Fprintln(out, 0)
			continue
		}
		q := make([]int, 0, n)
		for i, d := range d {
			if d == 1 {
				q = append(q, i)
			}
		}
		for ; k > 0 && len(q) > 0; k-- {
			sz := len(q)
			for _, v := range q {
				for _, w := range g[v] {
					if d[w]--; d[w] == 1 {
						q = append(q, w)
					}
				}
			}
			q = q[sz:]
		}
		Fprintln(out, cap(q))
	}
}

//func main() { CF1593E(os.Stdin, os.Stdout) }
