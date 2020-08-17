package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1316D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		raw := in.Bytes()
		if raw[0] == '-' {
			return -1
		}
		for _, b := range raw {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	ends := make([][][2]int, n)
	for i := range ends {
		ends[i] = make([][2]int, n)
		for j := range ends[i] {
			if x, y := read(), read(); x > 0 {
				ends[i][j] = [2]int{x - 1, y - 1}
			} else {
				ends[i][j] = [2]int{-1, -1}
			}
		}
	}

	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, n)
	}
	valid := func(x, y int) bool { return x >= 0 && x < n && y >= 0 && y < n }
	dir4 := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	const dirMark = "LURD"
	var tar [2]int
	var makePath func(x, y int)
	makePath = func(x, y int) {
		for i, d := range dir4 {
			if xx, yy := x+d[0], y+d[1]; valid(xx, yy) && ans[xx][yy] == 0 && ends[xx][yy] == tar {
				ans[xx][yy] = dirMark[i]
				makePath(xx, yy)
			}
		}
	}
	for i, row := range ends {
		for j, ed := range row {
			if ans[i][j] > 0 {
				continue
			}
			if x, y := ed[0], ed[1]; x == -1 {
				tar = [2]int{-1, -1}
				if valid(i, j+1) && ends[i][j+1] == tar {
					ans[i][j] = 'R'
					ans[i][j+1] = 'L'
					makePath(i, j)
					makePath(i, j+1)
				} else if valid(i+1, j) && ends[i+1][j] == tar {
					ans[i][j] = 'D'
					ans[i+1][j] = 'U'
					makePath(i, j)
					makePath(i+1, j)
				} else {
					Fprint(out, "INVALID")
					return
				}
			} else {
				if ans[x][y] > 0 || ends[x][y] != ed {
					Fprint(out, "INVALID")
					return
				}
				ans[x][y] = 'X'
				tar = ed
				makePath(x, y)
				if ans[i][j] == 0 {
					Fprint(out, "INVALID")
					return
				}
			}
		}
	}
	Fprintln(out, "VALID")
	for _, row := range ans {
		Fprintf(out, "%s\n", row)
	}
}

//func main() { CF1316D(os.Stdin, os.Stdout) }
