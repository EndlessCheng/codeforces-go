package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1621B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r, c int
	for Fscan(in, &T); T > 0; T-- {
		ll, rr, lc, rc := int(2e9), 0, 0, 0
		sl, sr, sc := int(2e9), 0, 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &l, &r, &c)
			if l < sl || l == sl && r > sr || l == sl && r == sr && c < sc {
				sl, sr, sc = l, r, c
			}
			if l < ll || l == ll && c < lc {
				ll, lc = l, c
			}
			if r > rr || r == rr && c < rc {
				rr, rc = r, c
			}
			res := lc + rc
			if sl == ll && sr == rr && sc < res {
				res = sc
			}
			Fprintln(out, res)
		}
	}
}

//func main() { CF1621B(os.Stdin, os.Stdout) }
