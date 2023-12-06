package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1907D(_r io.Reader, _w io.Writer) {
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
		ans := sort.Search(1e9, func(k int) bool {
			l, r := 0, k
			for _, p := range a {
				l = max(l, p.l)
				r = min(r, p.r)
				if l > r {
					return false
				}
				l -= k
				r += k
			}
			return true
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1907D(os.Stdin, os.Stdout) }
