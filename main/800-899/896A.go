package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF896A(in io.Reader, out io.Writer) {
	const (
		s  = `What are you doing at the end of the world? Are you busy? Will you save us?`
		p0 = `What are you doing while sending "`
		p1 = `"? Are you busy? Will you send "`
		p2 = `"?`
		l0 = int64(len(p0))
		l1 = int64(len(p1))
		l2 = int64(len(p2))
	)
	sz := []int64{int64(len(s))}
	for sz[len(sz)-1] < 1e18 {
		sz = append(sz, sz[len(sz)-1]*2+l0+l1+l2)
	}

	var T, n int
	var k int64
	f := func() byte {
		k--
		for {
			if n < len(sz) && k >= sz[n] {
				return '.'
			}
			if n == 0 {
				return s[k]
			}
			if k < l0 {
				return p0[k]
			}
			k -= l0
			n--
			if n >= len(sz) || k < sz[n] {
				continue
			}
			k -= sz[n]
			if k < l1 {
				return p1[k]
			}
			k -= l1
			if k < sz[n] {
				continue
			}
			k -= sz[n]
			return p2[k]
		}
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		Fprintf(out, "%c", f())
	}
}

//func main() { CF896A(os.Stdin, os.Stdout) }
