package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1539D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	a := make([]struct{ c, lim int64 }, n)
	for i := range a {
		Fscan(in, &a[i].c, &a[i].lim)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].lim < a[j].lim })

	ans := int64(0)
	for l, r, c := 0, n-1, int64(0); l <= r; {
		if a[l].lim > c {
			v := min(a[l].lim-c, a[r].c)
			ans += v * 2
			c += v
			if a[r].c -= v; a[r].c == 0 {
				r--
			}
		} else {
			ans += a[l].c
			c += a[l].c
			l++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1539D(os.Stdin, os.Stdout) }
