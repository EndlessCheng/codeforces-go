package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1626C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
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

		sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })
		ans := int64(0)
		l0, maxR := a[0].l, a[0].r
		for _, p := range a[1:] {
			l, r := p.l, p.r
			if l > maxR {
				sz := int64(maxR - l0 + 1)
				ans += (sz + 1) * sz / 2
				l0 = l
			}
			maxR = max(maxR, r)
		}
		sz := int64(maxR - l0 + 1)
		ans += (sz + 1) * sz / 2
		Fprintln(out, ans)
	}
}

//func main() { CF1626C(os.Stdin, os.Stdout) }
