package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1738C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	dp := [101][101][2][2]int8{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = [2][2]int8{{-1, -1}, {-1, -1}}
		}
	}
	var f func(c0, c1, s, who int8) int8
	f = func(c0, c1, s, who int8) (res int8) {
		if c0 == 0 && c1 == 0 {
			return s ^ who ^ 1
		}
		dv := &dp[c0][c1][s][who]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if c0 > 0 && f(c0-1, c1, s, who^1) == 0 || c1 > 0 && f(c0, c1-1, s^who^1, who^1) == 0 {
			return 1
		}
		return
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		cnt := [2]int8{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			cnt[v&1]++
		}
		if f(cnt[0], cnt[1], 0, 0) == 1 {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { CF1738C(os.Stdin, os.Stdout) }
