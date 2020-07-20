package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1385F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k, v, w int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if k == 1 {
			Fprintln(out, n-1)
			continue
		}

		ans := 0
		l := make([]int, n)
		used := make([]bool, n)
		q := []int{}
		for v, vs := range g {
			for _, w := range vs {
				if len(g[w]) == 1 {
					l[v]++
					used[w] = true
				}
			}
			if len(vs) > 1 && l[v]+1 >= len(vs) && l[v]%k == 0 {
				q = append(q, v)
			}
		}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			ans += l[v] / k
			l[v] = 0
			used[v] = true
			for _, w := range g[v] {
				if !used[w] {
					l[w]++
					if l[w]+1 >= len(g[w]) && l[w]%k == 0 {
						q = append(q, w)
					}
					break
				}
			}
		}
		for _, c := range l {
			ans += c / k
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1385F(os.Stdin, os.Stdout) }
