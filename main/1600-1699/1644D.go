package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1644D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var T, n, m, k, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k, &q)
		a := make([][2]int, q)
		for i := range a {
			Fscan(in, &a[i][0], &a[i][1])
		}
		ans := int64(1)
		x, y := map[int]bool{}, map[int]bool{}
		for i := q - 1; i >= 0 && len(x) < n && len(y) < m; i-- {
			r, c := a[i][0], a[i][1]
			if !x[r] || !y[c] {
				ans = ans * int64(k) % mod
				x[r] = true
				y[c] = true
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1644D(os.Stdin, os.Stdout) }
