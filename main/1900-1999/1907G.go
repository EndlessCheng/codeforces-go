package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1907G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		g := make([]int, n)
		deg := make([]int, n)
		for i := range g {
			Fscan(in, &g[i])
			g[i]--
			deg[g[i]]++
		}

		ans := []int{}
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			w := g[v]
			if s[v] == '1' {
				ans = append(ans, v)
				s[w] ^= 1
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}

		for i, d := range deg {
			if d == 0 {
				continue
			}
			ring := []int{i}
			p1 := []int{}
			if s[i] == '1' {
				p1 = append(p1, 0)
			}
			for v := g[i]; v != i; v = g[v] {
				deg[v] = 0
				if s[v] == '1' {
					p1 = append(p1, len(ring))
				}
				ring = append(ring, v)
			}
			if len(p1) == 0 {
				continue
			}
			if len(p1)%2 > 0 {
				Fprintln(out, -1)
				continue o
			}

			cnt := 0
			for i := 1; i < len(p1); i += 2 {
				cnt += p1[i] - p1[i-1]
			}
			idx := p1[0]
			if cnt > len(ring)-cnt {
				idx = p1[1]
			}

			ring = append(ring[idx:], ring[:idx]...)
			for j, v := range ring {
				if s[v] == '1' {
					ans = append(ans, v)
					s[ring[(j+1)%len(ring)]] ^= 1
				}
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1907G(os.Stdin, os.Stdout) }
