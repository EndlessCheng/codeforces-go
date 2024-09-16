package main

import (
	. "fmt"
	"io"
)

func cf1948C(in io.Reader, out io.Writer) {
	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var T, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		g := [2]string{}
		Fscan(in, &g[0], &g[1])
		vis := [2][]bool{make([]bool, m), make([]bool, m)}
		var dfs func(int, int) bool
		dfs = func(i, j int) bool {
			if i == 1 && j == m-1 {
				return true
			}
			vis[i][j] = true
			for _, d := range dir4 {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < 2 && 0 <= y && y < m {
					if g[x][y] == '<' {
						y--
					} else {
						y++
					}
					if !vis[x][y] && dfs(x, y) {
						return true
					}
				}
			}
			return false
		}
		if dfs(0, 0) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1948C(bufio.NewReader(os.Stdin), os.Stdout) }
