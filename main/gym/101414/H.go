package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func runH(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m, n, v int
	for {
		if Fscanln(in, &m, &n); m == 0 {
			break
		}
		cnt := make([][]int, 1<<m) // [猜集合][正确集合]
		for i := range cnt {
			cnt[i] = make([]int, 1<<m)
		}
		for ; n > 0; n-- {
			Fscanf(in, "%b\n", &v)
			for s, c := range cnt {
				c[s&v]++
			}
		}
		dp := make([][]int, 1<<m) // [猜集合][正确集合]
		for i := range dp {
			dp[i] = make([]int, 1<<m)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(s, t int) (res int) {
			if cnt[s][t] < 2 {
				return
			}
			if cnt[s][t] == 2 {
				return 1
			}
			dv := &dp[s][t]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
			res = 1e9
			for cs, lb := s^(1<<m-1), 0; cs > 0; cs ^= lb {
				lb = cs & -cs
				if s2, t2 := s|lb, t|lb; cnt[s2][t2] > 0 && cnt[s2][t] > 0 {
					res = min(res, 1+max(f(s2, t2), f(s2, t)))
				}
			}
			return
		}
		Fprintln(out, f(0, 0))
	}
}

func main() { r, _ := os.Open("H.in"); runH(r, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
