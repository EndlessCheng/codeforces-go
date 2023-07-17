package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1841D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
		}
		sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })

		ans := n
		r1, r2 := -1, -1
		for _, p := range a {
			if r2 < p.l {
				if r1 < p.l {
					r1 = p.r
				} else {
					ans -= 2
					r2 = p.r
					r1 = -1
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1841D(os.Stdin, os.Stdout) }
