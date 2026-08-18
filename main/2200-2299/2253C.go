package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2253C(in io.Reader, out io.Writer) {
	var T, n, m, x, y, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &x, &y)
		tp := make([]byte, n+m+1)
		for range x {
			Fscan(in, &v)
			tp[v] = 1
		}
		for range y {
			Fscan(in, &v)
			tp[v] |= 2
		}

		var ans, row, col, common int
		for v := n + m; v > 0 && row+col+common < n+m-1; v-- {
			t := tp[v]
			if t == 1 {
				if row < n {
					ans += v
					row++
				}
			} else if t == 2 {
				if col < m {
					ans += v
					col++
				}
			} else if t == 3 {
				ans += v
				common++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2253C(bufio.NewReader(os.Stdin), os.Stdout) }
