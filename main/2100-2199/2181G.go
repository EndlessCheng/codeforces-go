package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2181G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = s[i-1] + a[i]
		}

		ans := (s[n]/2 + n - 2) / (n - 1)
		if n%2 > 0 {
			x := 0
			for i := 1; i <= n; i++ {
				if i%2 > 0 {
					x += a[i]
				} else {
					x -= a[i]
				}
			}
			x /= 2
			ans = max(ans, x)
			for i := 1; i <= n; i++ {
				x = a[i] - x
				ans = max(ans, x)
			}
			Fprintln(out, ans)
			continue
		}

		l, r := 0, int(1e18)
		p, q := 0, int(-1e18)
		x := 0
		for i := 1; i <= n; i++ {
			x = a[i] - x
			if i%2 > 0 {
				r = min(r, x)
				q = max(q, x)
			} else {
				l = max(l, -x)
				p = max(p, x)
			}
		}

		mid := min(max((q-p)/2, l), r)
		Fprintln(out, max(ans, p+mid, q-mid))
	}
}

//func main() { cf2181G(bufio.NewReader(os.Stdin), os.Stdout) }
