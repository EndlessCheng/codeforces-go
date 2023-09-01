package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1859D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ l, r int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].l, &n, &n, &a[i].r)
		}
		sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })
		b := a[:0]
		l0, maxR := a[0].l, a[0].r
		for _, p := range a[1:] {
			if p.l > maxR {
				b = append(b, pair{l0, maxR})
				l0 = p.l
			}
			maxR = max(maxR, p.r)
		}
		b = append(b, pair{l0, maxR})
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &x)
			i := sort.Search(len(b), func(i int) bool { return b[i].r >= x })
			if i < len(b) && b[i].l <= x {
				x = b[i].r
			}
			Fprint(out, x, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1859D(os.Stdin, os.Stdout) }
