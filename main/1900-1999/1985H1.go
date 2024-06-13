package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1985H1(in io.Reader, out io.Writer) {
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var T, n, m, minX, minY, maxX, maxY, sz int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		var dfs func(int, int)
		dfs = func(i, j int) {
			minX = min(minX, i)
			maxX = max(maxX, i)
			minY = min(minY, j)
			maxY = max(maxY, j)
			sz++
			a[i][j] = 0
			for _, d := range dir4 {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < m && a[x][y] == '#' {
					dfs(x, y)
				}
			}
		}
		cr := make([]int, n)
		cc := make([]int, m)
		dr := make([]int, n+2)
		dc := make([]int, m+2)
		for i, row := range a {
			for j, b := range row {
				if b == '#' {
					minX, minY, maxX, maxY, sz = n, m, 0, 0, 0
					dfs(i, j)
					minX = max(minX-1, 0)
					minY = max(minY-1, 0)
					dr[minX] += sz
					dr[maxX+2] -= sz
					dc[minY] += sz
					dc[maxY+2] -= sz
				} else if b == '.' {
					cr[i]++
					cc[j]++
				}
			}
		}
		ans := 0
		s := 0
		for i, d := range dr[:n] {
			s += d
			ans = max(ans, cr[i]+s)
		}
		s = 0
		for j, d := range dc[:m] {
			s += d
			ans = max(ans, cc[j]+s)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1985H1(bufio.NewReader(os.Stdin), os.Stdout) }
