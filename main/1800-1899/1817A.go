package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1817A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	g1 := []int{}
	type pair struct{ l, r int }
	g2 := []pair{}
	for i := 0; i < n; {
		st := i
		for i++; i < n && a[i-1] >= a[i]; i++ {
		}
		if i-st == 1 {
			g1 = append(g1, st)
		} else {
			g2 = append(g2, pair{st, i})
		}
	}

	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		l--
		ans := sort.SearchInts(g1, r) - sort.SearchInts(g1, l)
		li := sort.Search(len(g2), func(i int) bool { return g2[i].r > l })
		if li < len(g2) && r > g2[li].l {
			ri := sort.Search(len(g2), func(i int) bool { return g2[i].l >= r }) - 1
			l1 := max(l, g2[li].l)
			r1 := min(r, g2[ri].r)
			if li == ri {
				ans += min(r1-l1, 2)
			} else {
				ans += min(g2[li].r-l1, 2) + (ri-li-1)*2 + min(r1-g2[ri].l, 2)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1817A(os.Stdin, os.Stdout) }
