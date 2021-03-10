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

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	dist := make([][]int, n)
	for i := range dist {
		d := make([]int, n)
		for j := range d {
			d[j] = -1
		}
		d[i] = 0
		q := []int{i}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			for _, w := range g[v] {
				if d[w] < 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		dist[i] = d
	}
	for i, di := range dist {
		for j, dij := range di {
			seen := make([]bool, n)
			mul := int64(1)
			for k, dik := range di {
				if k == i || k == j {
					continue
				}
				if dik+dist[j][k] == dij { // i-j 最短路不唯一
					if seen[dik] {
						mul = 0
						break
					}
					seen[dik] = true
				} else {
					c := 0
					for _, w := range g[k] {
						if di[w]+1 == dik && dist[j][w]+1 == dist[j][k] {
							c++
						}
					}
					mul = mul * int64(c) % 998244353
				}
			}
			Fprint(out, mul, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1495D(os.Stdin, os.Stdout) }
