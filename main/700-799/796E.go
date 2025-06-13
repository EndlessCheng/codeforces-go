package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf796E(in io.Reader, out io.Writer) {
	var n, p, k, m, v, ans int
	Fscan(in, &n, &p, &k)
	done := make([][2]bool, n+1)
	for i := range 2 {
		Fscan(in, &m)
		for range m {
			Fscan(in, &v)
			done[v][i] = true
		}
	}
	for _, p := range done {
		if p[0] || p[1] {
			ans++
		}
	}
	if p/2*k >= n {
		Fprint(out, ans)
		return
	}

	dp := make([][][][2]int16, k+1)
	for i := range dp {
		dp[i] = make([][][2]int16, p+1)
		for j := range dp[i] {
			dp[i][j] = make([][2]int16, n+1)
			for l := range dp[i][j] {
				dp[i][j][l] = [2]int16{-1, -1}
			}
		}
	}
	var f func(int, int, int, int) int16
	f = func(i, p, ex, who int) (miss int16) {
		if i <= 0 {
			return
		}
		t := &dp[ex][p][i][who]
		if *t >= 0 {
			return *t
		}

		miss = f(i-1, p, max(ex-1, 0), who)
		d := done[i]
		if ex == 0 && (d[0] || d[1]) || ex > 0 && !d[who] && d[who^1] {
			miss++
		}

		if p > 0 {
			if ex == 0 {
				miss = min(miss, f(i, p-1, k, 0), f(i, p-1, k, 1))
			} else {
				miss = min(miss, f(i-ex, p-1, k-ex, who^1))
			}
		}

		*t = miss
		return
	}
	Fprint(out, ans-int(f(n, p, 0, 0)))
}

//func main() { cf796E(bufio.NewReader(os.Stdin), os.Stdout) }
