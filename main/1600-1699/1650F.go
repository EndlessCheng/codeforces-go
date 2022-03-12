package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1650F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, e, t, p int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		type pair struct{ t, p, i int }
		gs := make([][]pair, n)
		for i := 1; i <= m; i++ {
			Fscan(in, &e, &t, &p)
			gs[e-1] = append(gs[e-1], pair{t, p, i})
		}

		ans := []interface{}{}
		t := 0
		for k, g := range gs {
			dp := [101]int{}
			for i := range dp {
				dp[i] = 1e9 + 1
			}
			dp[0] = 0
			from := make([][101]int, len(g))
			for i, p := range g {
				for j := 100; j > 0; j-- {
					k := j - p.p
					if k < 0 {
						k = 0
					}
					if v := dp[k] + p.t; v < dp[j] {
						dp[j] = v
						from[i][j] = k
					} else {
						from[i][j] = j
					}
				}
			}
			t += dp[100]
			if t > a[k] {
				Fprintln(out, -1)
				continue o
			}
			for i, j := len(g)-1, 100; i >= 0; i-- {
				if from[i][j] != j {
					ans = append(ans, g[i].i)
					j = from[i][j]
				}
			}
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1650F(os.Stdin, os.Stdout) }
