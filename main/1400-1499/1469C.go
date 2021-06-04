package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1469C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, k, h, l int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &l)
		s := "YES"
		r := l
		for ; n > 2; n-- {
			Fscan(in, &h)
			l = max(l+1-k, h)
			r = min(r, h) + k - 1
			if l > r {
				s = "NO"
			}
		}
		Fscan(in, &h)
		if l-k >= h || r+k <= h {
			s = "NO"
		}
		Fprintln(out, s)
	}
}

//func main() { CF1469C(os.Stdin, os.Stdout) }
