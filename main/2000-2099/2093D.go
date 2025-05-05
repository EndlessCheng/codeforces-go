package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 注：也可以用二进制做

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
				u, d, l, r := 0, 1<<n, 0, 1<<n
				for s := 1 << (n*2 - 2); s > 0; s >>= 2 {
					xm, ym := (u+d)/2, (l+r)/2
					if x < xm && y < ym {
						d, r = xm, ym
					} else if x >= xm && y >= ym {
						ans += s
						u, l = xm, ym
					} else if x >= xm && y < ym {
						ans += s * 2
						u, r = xm, ym
					} else {
						ans += s * 3
						d, l = xm, ym
					}
				}
				Fprintln(out, ans)
			} else {
				u, d, l, r := 0, 1<<n, 0, 1<<n
				for s := 1 << (n*2 - 2); s > 0; s >>= 2 {
					xm, ym := (u+d)/2, (l+r)/2
					if x < s {
						d, r = xm, ym
					} else if x < s*2 {
						x -= s
						u, l = xm, ym
					} else if x < s*3 {
						x -= s * 2
						u, r = xm, ym
					} else {
						x -= s * 3
						d, l = xm, ym
					}
				}
				Fprintln(out, u+1, l+1)
			}
		}
	}
}

func main() { cf2093D(bufio.NewReader(os.Stdin), os.Stdout) }
