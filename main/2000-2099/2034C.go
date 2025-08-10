package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2034C(in io.Reader, out io.Writer) {
	type pair struct{ x, y int }
	dirC := [...]pair{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	dir4 := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		dp := make([][]int8, n)
		for i := range dp {
			dp[i] = make([]int8, m)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int8
		f = func(i, j int) (res int8) {
			if i < 0 || i >= n || j < 0 || j >= m {
				return 0
			}
			p := &dp[i][j]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
			*p = 1
			b := a[i][j]
			if b != '?' {
				d := dirC[b]
				return f(i+d.x, j+d.y)
			}
			for _, d := range dir4 {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < m && (a[x][y] == '?' || f(x, y) > 0) {
					return 1
				}
			}
			return 0
		}

		ans := 0
		for i := range n {
			for j := range m {
				ans += int(f(i, j))
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2034C(bufio.NewReader(os.Stdin), os.Stdout) }
