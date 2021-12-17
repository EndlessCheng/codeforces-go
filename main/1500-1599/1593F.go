package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1593F(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, ma, mb int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &ma, &mb, &s)
		minD, ans := n, "-1"
		t := make([]byte, n)
		vis := [40][40][40][41]bool{}
		var f func(i, x, y, c int)
		f = func(i, x, y, c int) {
			if i == n {
				if x == 0 && y == 0 && abs(n-c*2) < minD {
					minD = abs(n - c*2)
					ans = string(t)
				}
				return
			}
			if vis[i][x][y][c] {
				return
			}
			vis[i][x][y][c] = true
			t[i] = 'R'
			f(i+1, (x*10+int(s[i]&15))%ma, y, c+1)
			t[i] = 'B'
			f(i+1, x, (y*10+int(s[i]&15))%mb, c)
		}
		f(0, 0, 0, 0)
		Fprintln(out, ans)
	}
}

//func main() { CF1593F(os.Stdin, os.Stdout) }
