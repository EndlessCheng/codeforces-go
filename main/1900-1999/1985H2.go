package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1985H2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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
		d2 := make([][]int, n+3)
		for i := range d2 {
			d2[i] = make([]int, m+3)
		}
		update := func(r1, c1, r2, c2, val int) {
			d2[r1+1][c1+1] += val
			d2[r1+1][c2+2] -= val
			d2[r2+2][c1+1] -= val
			d2[r2+2][c2+2] += val
		}
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
					update(minX, minY, maxX+1, maxY+1, sz)
				} else if b == '.' {
					cr[i]++
					cc[j]++
				}
			}
		}
		for i := 1; i < n; i++ {
			dr[i] += dr[i-1]
		}
		for j := 1; j < m; j++ {
			dc[j] += dc[j-1]
		}
		ans := 0
		for i, row := range a {
			for j, b := range row {
				d2[i+1][j+1] += d2[i+1][j] + d2[i][j+1] - d2[i][j]
				res := cr[i] + cc[j] + dr[i] + dc[j] - d2[i+1][j+1]
				if b == '.' {
					res--
				}
				ans = max(ans, res)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1985H2(bufio.NewReader(os.Stdin), os.Stdout) }
