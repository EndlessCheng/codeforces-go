package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1843E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		qs := make([]struct{ l, r int }, m)
		for i := range qs {
			Fscan(in, &qs[i].l, &qs[i].r)
		}
		Fscan(in, &q)
		a := make([]int, q)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := sort.Search(q+1, func(x int) bool {
			s := make([]int, n+1)
			for _, i := range a[:x] {
				s[i] = 1
			}
			for i := 2; i <= n; i++ {
				s[i] += s[i-1]
			}
			for _, q := range qs {
				if (s[q.r]-s[q.l-1])*2 > q.r-q.l+1 {
					return true
				}
			}
			return false
		})
		if ans > q {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1843E(os.Stdin, os.Stdout) }
