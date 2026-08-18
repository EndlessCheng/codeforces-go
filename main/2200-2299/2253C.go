package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func cf2253C(in io.Reader, out io.Writer) {
	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &x, &y)
		sz := n + m + 1
		a := make([]int, sz)
		for i := sz - x; i < sz; i++ {
			Fscan(in, &a[i])
		}
		b := make([]int, sz)
		for i := sz - y; i < sz; i++ {
			Fscan(in, &b[i])
		}

		f := func(n, m int) (res int) {
			has := map[int]bool{}
			same := 0
			i, j := 1, 1
			for i <= n || j <= m {
				v, w := 0, 0
				if i <= n {
					v = a[sz-i]
				}
				if j <= m {
					w = b[sz-j]
				}
				if v == 0 && w == 0 {
					break
				}
				has[max(v, w)] = true
				if v > w {
					res += v
					i++
				} else if v < w {
					res += w
					j++
				} else {
					res += v
					same++
					i++
					j++
				}
			}
			for range same {
				for has[a[sz-i]] {
					i++
				}
				v := a[sz-i]
				for has[b[sz-j]] {
					j++
				}
				w := b[sz-j]
				if v == 0 && w == 0 {
					break
				}
				has[max(v, w)] = true
				if v > w {
					res += v
					i++
				} else {
					res += w
					j++
				}
			}
			return
		}
		Fprintln(out, max(f(n, m-1), f(n-1, m)))
	}
}

func main() { cf2253C(bufio.NewReader(os.Stdin), os.Stdout) }
