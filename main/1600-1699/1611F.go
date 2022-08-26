package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1611F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		a := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			a[i] += a[i-1]
		}
		ll, rr := -1, -1
		for l, r := 0, 0; r < n; l++ {
			for r < n && a[r+1]-a[l] >= -s {
				r++
			}
			if r-l > rr-ll {
				ll, rr = l, r
			}
		}
		if ll < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ll+1, rr)
		}
	}
}

//func main() { CF1611F(os.Stdin, os.Stdout) }
