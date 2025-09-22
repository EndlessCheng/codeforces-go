package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1687C(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
		}
		for i := 1; i <= n; i++ {
			var x int
			Fscan(in, &x)
			s[i] += s[i-1] - x
		}
		g := make([][]int, n+1)
		for range m {
			var l, r int
			Fscan(in, &l, &r)
			l--
			g[l] = append(g[l], r)
			g[r] = append(g[r], l)
		}
		if s[n] != 0 {
			Fprintln(out, "NO")
			continue
		}

		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		find := func(x int) int {
			rt := x
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[x] != rt {
				fa[x], x = rt, fa[x]
			}
			return rt
		}

		q := []int{}
		for i, s := range s[:n] {
			if s == 0 {
				fa[i] = i + 1
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			n--
			l := q[0]
			q = q[1:]
			for _, r := range g[l] {
				if s[r] != 0 {
					continue
				}
				l := l
				if l > r {
					l, r = r, l
				}
				to := find(r)
				for j := find(l); j < r; j = find(j + 1) {
					s[j] = 0
					fa[j] = to
					q = append(q, j)
				}
			}
		}
		if n == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1687C(bufio.NewReader(os.Stdin), os.Stdout) }
