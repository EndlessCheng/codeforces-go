package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1613E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		x, y := 0, -1
		g := make([][]byte, n)
		for i := range g {
			Fscan(in, &g[i])
			if y < 0 {
				x, y = i, bytes.IndexByte(g[i], 'L')
			}
		}
		for _, d := range dir4 {
			x, y := x+d.x, y+d.y
			if 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
			o:
				for {
					dir := pair{}
					for _, d := range dir4 {
						if x, y := x+d.x, y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
							if dir != (pair{}) {
								break o
							}
							dir = d
						}
					}
					g[x][y] = '+'
					if dir == (pair{}) {
						break
					}
					x += dir.x
					y += dir.y
				}
			}
		}
		for _, r := range g {
			Fprintf(out, "%s\n", r)
		}
	}
}

//func main() { CF1613E(os.Stdin, os.Stdout) }
