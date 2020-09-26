package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF543B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, s1, t1, l1, s2, t2, l2 int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	d := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		d[i] = make([]int, n+1)
		for j := range d[i] {
			d[i][j] = -1
		}
		d[i][i] = 0
		q := []int{i}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			for _, w := range g[v] {
				if d[i][w] < 0 {
					d[i][w] = d[i][v] + 1
					q = append(q, w)
				}
			}
		}
	}

	Fscan(in, &s1, &t1, &l1, &s2, &t2, &l2)
	if d[s1][t1] > l1 || d[s2][t2] > l2 {
		Fprint(out, -1)
		return
	}
	ans := m - d[s1][t1] - d[s2][t2]
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if d[i][s1]+d[i][j]+d[j][t1] <= l1 && d[i][s2]+d[i][j]+d[j][t2] <= l2 {
				if v := m - d[i][j] - d[i][s1] - d[i][s2] - d[j][t1] - d[j][t2]; v > ans {
					ans = v
				}
			}
			if d[i][s1]+d[i][j]+d[j][t1] <= l1 && d[i][t2]+d[i][j]+d[j][s2] <= l2 {
				if v := m - d[i][j] - d[i][s1] - d[i][t2] - d[j][t1] - d[j][s2]; v > ans {
					ans = v
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF543B(os.Stdin, os.Stdout) }
