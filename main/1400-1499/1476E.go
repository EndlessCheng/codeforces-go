package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1476E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, v int
	Fscan(in, &n, &m, &k)
	mp := make(map[string]int, n)
	for i := 1; i <= n; i++ {
		var s string
		Fscan(in, &s)
		mp[s] = i
	}

	var s []byte
	var ids []int
	var f func(int)
	f = func(p int) {
		if p == k {
			if id := mp[string(s)]; id > 0 {
				ids = append(ids, id-1)
			}
			return
		}
		f(p + 1)
		tmp := s[p]
		s[p] = '_'
		f(p + 1)
		s[p] = tmp
	}

	g := make([][]int, n)
	deg := make([]int, n)
o:
	for i := 0; i < m; i++ {
		Fscan(in, &s, &v)
		v--
		ids = []int{}
		f(0)
		for _, id := range ids {
			if id == v {
				for _, w := range ids {
					if w != v {
						g[v] = append(g[v], w)
						deg[w]++
					}
				}
				continue o
			}
		}
		Fprintln(out, "NO")
		return
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
		ans = append(ans, v)
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	if len(ans) < n {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF1476E(os.Stdin, os.Stdout) }
