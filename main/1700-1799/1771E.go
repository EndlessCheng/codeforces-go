package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1771E(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	s := make([]string, n+1)
	a := make([][]int, n+1)
	v := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, m+1)
		v[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		for j := 1; j <= m; j++ {
			a[i][j] = b2i71(s[i][j-1] == 'm') + 2*b2i71(s[i][j-1] == '#')
			v[i][j] = v[i][j-1] + a[i][j]
		}
	}

	for i := 1; i <= m; i++ {
		for j := i + 2; j <= m; j++ {
			l, r, p0, p1, cnt := 1, 1, 0, 0, 0
			for r <= n {
				if a[r][i]+a[r][j] > 1 {
					r++
					l = r
					cnt = 0
					p0 = 0
					p1 = 0
					continue
				}
				cnt += a[r][i] + a[r][j]
				for cnt >= 2 && l <= r {
					cnt -= a[l][i] + a[l][j]
					l++
				}
				if r-l > 1 && (cnt < 2 && p0 > l || cnt == 0 && p1 > l) {
					ans = max(ans, j-i+1+2*(r-l))
				}
				if v[r][j-1]-v[r][i] == 1 {
					p1 = r
				}
				if v[r][j-1]-v[r][i] == 0 {
					p0 = r
				}
				r++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1771E(bufio.NewReader(os.Stdin), os.Stdout) }

func b2i71(b bool) int {
	if b {
		return 1
	}
	return 0
}
