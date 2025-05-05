package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2093D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, x, y int
	var op string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		for range q {
			Fscan(in, &op, &x)
			x--
			if op[0] == '-' {
				Fscan(in, &y)
				y--
				ans := 1
				for b := 1; b < 1<<n; b <<= 1 {
					if x&b > 0 && y&b > 0 {
						ans += b * b
					} else if x&b > 0 {
						ans += b * b * 2
					} else if y&b > 0 {
						ans += b * b * 3
					}
				}
				Fprintln(out, ans)
			} else {
				r, c := 0, 0
				for i := range n {
					if x%4 == 1 || x%4 == 2 {
						r |= 1 << i
					}
					if x%2 > 0 {
						c |= 1 << i
					}
					x /= 4
				}
				Fprintln(out, r+1, c+1)
			}
		}
	}
}

//func main() { cf2093D(bufio.NewReader(os.Stdin), os.Stdout) }
