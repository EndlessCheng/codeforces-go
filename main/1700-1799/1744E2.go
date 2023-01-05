package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1744E2(in io.Reader, out io.Writer) {
	f := func(a int) (ds []int) {
		for d := 1; d*d <= a; d++ {
			if a%d == 0 {
				ds = append(ds, d)
				if d*d < a {
					ds = append(ds, a/d)
				}
			}
		}
		sort.Ints(ds)
		return
	}

	var T, a, b, c, d int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &d)
		ds := f(b)
		for _, v := range f(a) {
			vv := a / v
			for _, w := range ds {
				if w > c/v {
					break
				}
				ww := b / w
				if ww > d/vv {
					continue
				}
				x, y := c-c%(v*w), d-d%(vv*ww)
				if x > a && y > b {
					Fprintln(out, x, y)
					continue o
				}
			}
		}
		Fprintln(out, -1, -1)
	}
}

//func main() { CF1744E2(os.Stdin, os.Stdout) }
