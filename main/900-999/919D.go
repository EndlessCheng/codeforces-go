package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF919D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, c int
	var s []byte
	Fscan(in, &n, &m, &s)
	g := make([][]int, n)
	d := make([]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		d[w]++
	}

	ans := 1
	dp := make([][26]int, n)
	q := []int{}
	for i, v := range d {
		if v == 0 {
			q = append(q, i)
			dp[i][s[i]-'a'] = 1
		}
	}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		c++
		for _, w := range g[v] {
			for i, dv := range dp[w] {
				x := 0
				if byte(i) == s[w]-'a' {
					x = 1
				}
				if dv < dp[v][i]+x {
					dp[w][i] = dp[v][i] + x
					if dp[w][i] > ans {
						ans = dp[w][i]
					}
				}
			}
			d[w]--
			if d[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if c < n {
		ans = -1
	}
	Fprint(_w, ans)
}

//func main() { CF919D(os.Stdin, os.Stdout) }
