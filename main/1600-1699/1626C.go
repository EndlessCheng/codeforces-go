package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1626C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].r)
		}
		for i := range a {
			Fscan(in, &a[i].l)
			a[i].l = a[i].r - a[i].l + 1
		}

		ans := int64(0)
		r0, minL := a[n-1].r, a[n-1].l
		for i := n-2; i >= 0; i-- {
			l, r := a[i].l, a[i].r
			if r < minL {
				sz := int64(r0 - minL + 1)
				ans += (sz + 1) * sz / 2
				r0 = r
			}
			minL = min(minL, l)
		}
		sz := int64(r0 - minL + 1)
		ans += (sz + 1) * sz / 2
		Fprintln(out, ans)
	}
}

//func main() { CF1626C(os.Stdin, os.Stdout) }
