package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf264C(in io.Reader, out io.Writer) {
	var n, q, a, b, mx int
	Fscan(in, &n, &q)
	ps := make([]struct{ v, c int }, n)
	for i := range ps {
		Fscan(in, &ps[i].v)
	}
	for i := range ps {
		Fscan(in, &ps[i].c)
	}
	f := make([]int, n+1)
	for ; q > 0; q-- {
		Fscan(in, &a, &b)
		for i := range f {
			f[i] = -1e18
		}
		var mx1, mx2, mxC int
		for _, p := range ps {
			c := p.c
			if c != mxC {
				mx = mx1
			} else {
				mx = mx2
			}
			f[c] = max(f[c]+max(p.v*a, 0), mx+p.v*b)
			if f[c] > mx1 {
				if c != mxC {
					mx2 = mx1
					mxC = c
				}
				mx1 = f[c]
			} else if c != mxC && f[c] > mx2 {
				mx2 = f[c]
			}
		}
		Fprintln(out, mx1)
	}
}

//func main() { cf264C(bufio.NewReader(os.Stdin), os.Stdout) }
