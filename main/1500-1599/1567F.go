package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1567F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n+2)
	ans := make([][]int, n+2)
	for i := range a {
		a[i] = make([]byte, m+2)
		ans[i] = make([]int, m+2)
	}
	for i := 1; i <= n; i++ {
		var s string
		Fscan(in, &s)
		for j := 1; j <= m; j++ {
			a[i][j] = s[j-1]
		}
	}

	f := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i][j] == 'X' {
				c := 0
				if a[i-1][j] == 'X' {
					f ^= 1
				}
				for _, d := range dir4 {
					if a[i+d.x][j+d.y] == '.' {
						c++
					}
				}
				if c&1 != 0 {
					Fprintln(out, "NO")
					return
				}
				ans[i][j] = c * 5 / 2
			} else {
				if j&1^f&1 != 0 {
					ans[i][j] = 1
				} else {
					ans[i][j] = 4
				}
			}
		}
	}

	Fprintln(out, "YES")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fprint(out, ans[i][j], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1567F(bufio.NewReader(os.Stdin), os.Stdout) }
