package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1495D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	dis := make([][]int, n)
	for i := range dis {
		d := make([]int, n)
		for j := range d {
			d[j] = -1
		}
		d[i] = 0
		q := []int{i}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, w := range g[v] {
				if d[w] < 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		dis[i] = d
	}

	seen := make([]bool, n)
	for i, di := range dis {
		for j, dij := range di {
			clear(seen)
			ans := 1
			for k, dik := range di {
				if k == i || k == j {
					continue
				}
				if dik+dis[j][k] == dij {
					if seen[dik] { // i-j 最短路不唯一
						ans = 0
						break
					}
					seen[dik] = true
				} else {
					c := 0
					for _, w := range g[k] {
						if di[w]+1 == dik && dis[j][w]+1 == dis[j][k] {
							c++
						}
					}
					ans = ans * c % mod
				}
			}
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1495D(os.Stdin, os.Stdout) }
