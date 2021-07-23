package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1163D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s, x, y []byte
	Fscan(in, &s, &x, &y)
	n, nx, ny := len(s), len(x), len(y)

	fail := func(s []byte) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	mx, my := fail(x), fail(y)
	update := func(b byte, cx, cy int) (int, int, int) {
		d := 0
		for cx > 0 && x[cx] != b {
			cx = mx[cx-1]
		}
		if x[cx] == b {
			cx++
		}
		if cx == nx {
			d++
			cx = mx[cx-1]
		}
		for cy > 0 && y[cy] != b {
			cy = my[cy-1]
		}
		if y[cy] == b {
			cy++
		}
		if cy == ny {
			d--
			cy = my[cy-1]
		}
		return d, cx, cy
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, nx)
		for j := range dp[i] {
			dp[i][j] = make([]int, ny)
			for k := range dp[i][j] {
				dp[i][j][k] = -1e9
			}
		}
	}
	var f func(p, cx, cy int) int
	f = func(p, cx, cy int) (res int) {
		if p == n {
			return
		}
		dv := &dp[p][cx][cy]
		if *dv > -1e9 {
			return *dv
		}
		defer func() { *dv = res }()
		b := s[p]
		if b != '*' {
			d, cx, cy := update(b, cx, cy)
			return d + f(p+1, cx, cy)
		}
		res = -1e9
		for b = 'a'; b <= 'z'; b++ {
			d, cx, cy := update(b, cx, cy)
			res = max(res, d+f(p+1, cx, cy))
		}
		return
	}
	Fprint(out, f(0, 0, 0))
}

//func main() { CF1163D(os.Stdin, os.Stdout) }
