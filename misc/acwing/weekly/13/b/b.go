package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	perm4 := [][4]int{{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1}, {1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0}, {2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0}, {3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0}}
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var T, n, m, sx, sy int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
			if j := bytes.IndexByte(a[i], 'S'); j >= 0 {
				sx, sy = i, j
			}
		}
		Fscan(in, &s)
		ans := 0
		for _, p := range perm4 {
			x, y := sx, sy
			for _, b := range s {
				d := dir4[p[b&3]]
				x += d.x
				y += d.y
				if x < 0 || x >= n || y < 0 || y >= m || a[x][y] == '#' {
					break
				}
				if a[x][y] == 'E' {
					ans++
					break
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
