package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF1811F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	f := func() bool {
		var n, m, v, w, c4, cur int
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for i := 0; i < m; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		k := int(math.Sqrt(float64(n)))
		if k < 3 || k*k != n || k*(k+1) != m {
			return false
		}
		for i, ws := range g {
			if len(ws) == 4 {
				c4++
				cur = i
			} else if len(ws) != 2 {
				return false
			}
		}
		if c4 != k {
			return false
		}
		vis := make([]bool, n)
		for !vis[cur] {
			vis[cur] = true
			for _, w := range g[cur] {
				if vis[w] {
					continue
				}
				if len(g[w]) == 4 {
					cur = w
					continue
				}
				cnt := 1
				for !vis[w] {
					vis[w] = true
					cnt++
					for _, w2 := range g[w] {
						if !vis[w2] {
							w = w2
							break
						}
					}
				}
				if cnt != k {
					return false
				}
				c4--
			}
		}
		return c4 == 0
	}
	var T int
	for Fscan(in, &T); T > 0; T-- {
		if f() {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1811F(os.Stdin, os.Stdout) }
