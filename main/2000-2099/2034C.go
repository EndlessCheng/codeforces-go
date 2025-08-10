package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func cf2034C(in io.Reader, out io.Writer) {
	type pair struct{ x, y int }
	dirC := [...]pair{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	rev := [...]byte{'L': 'R', 'R': 'L', 'U': 'D', 'D': 'U'}
	dir4 := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var T, n, m, ts int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0
		vis := make([][]int, n)
		for i := range vis {
			vis[i] = make([]int, m)
		}
		bad := make([][]bool, n)
		for i := range bad {
			bad[i] = make([]bool, m)
		}
		var canExit func(int, int) bool
		canExit = func(i, j int) bool {
			vis[i][j] = ts
			b := a[i][j]
			d := dirC[b]
			x, y := i+d.x, j+d.y
			if x < 0 || x >= n || y < 0 || y >= m {
				return true
			}
			if vis[x][y] > 0 && vis[x][y] < ts {
				b := bad[x][y]
				if b {
					ans++
				}
				bad[i][j] = b
				return !b
			}
			if vis[x][y] == ts {
				ans++
				bad[i][j] = true
				return false
			}
			if a[x][y] == '?' {
				ans += 2
				a[x][y] = rev[b]
				vis[x][y] = ts
				bad[x][y] = true
				bad[i][j] = true
				return false
			}
			if canExit(x, y) {
				return true
			}
			ans++
			bad[i][j] = true
			return false
		}
		for i, row := range a {
			for j, b := range row {
				if b != '?' && vis[i][j] == 0 {
					ts++
					canExit(i, j)
				}
			}
		}

		sz := 0
		found := false
		var dfs func(int, int)
		dfs = func(i, j int) {
			sz++
			a[i][j] = 0
			for _, d := range dir4 {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < m {
					if a[x][y] == '?' {
						dfs(x, y)
					} else if bad[x][y] {
						found = true
					}
				}
			}
		}
		for i, row := range a {
			for j, b := range row {
				if b == '?' {
					sz = 0
					found = false
					dfs(i, j)
					if found || sz > 1 {
						ans += sz
					}
				}
			}
		}

		Fprintln(out, ans)
	}
}

func main() { cf2034C(bufio.NewReader(os.Stdin), os.Stdout) }
