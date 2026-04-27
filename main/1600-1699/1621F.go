package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1621F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, a, b, c int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b, &c, &s)
		s = " " + s
		var ca, cb, cc, ce int
		for i := 1; i < n; i++ {
			if s[i] == '0' && s[i+1] == '0' {
				ca++
			} else if s[i] == '1' && s[i+1] == '1' {
				cb++
			}
		}

		t := min(ca, cb)
		ca -= t
		cb -= t
		ans := t * (a + b)
		if ca > 0 {
			ca--
			ans += a
		}
		if cb > 0 {
			cb--
			ans += b
		}

		if b > c {
			x := []int{}
			for i := 1; i <= n; i++ {
				if s[i] == '1' {
					continue
				}
				r := i
				for r < n && s[r+1] == '0' {
					r++
				}
				if i == 1 || r == n {
					ce++
				} else {
					x = append(x, r-i)
				}
				i = r
			}
			slices.Sort(x)
			for _, v := range x {
				if t >= v {
					t -= v
					cc++
				}
			}
			ans += (b - c) * (cc + min(cb, ce))
		}

		Fprintln(out, ans)
	}
}

//func main() { cf1621F(bufio.NewReader(os.Stdin), os.Stdout) }
