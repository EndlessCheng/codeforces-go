package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick42 []struct{ maxL, minR int }

func (f fenwick42) update(i, l, r int) {
	for ; i < len(f); i += i & -i {
		f[i].maxL = max(f[i].maxL, l)
		f[i].minR = min(f[i].minR, r)
	}
}

func (f fenwick42) pre(i int) int {
	maxL, minR := 0, int(2e9)
	for ; i > 0; i &= i - 1 {
		maxL = max(maxL, f[i].maxL)
		minR = min(minR, f[i].minR)
	}
	return minR - maxL
}

func cf2042D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type tuple struct{ l, r, i int }
		a := make([]tuple, n)
		ls := make([]int, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
			a[i].i = i
			ls[i] = a[i].l
		}
		slices.SortFunc(a, func(a, b tuple) int { return cmp.Or(b.r-a.r, a.l-b.l) })
		slices.Sort(ls)

		ans := make([]int, n)
		t := make(fenwick42, n+1)
		for i := range t {
			t[i].minR = 2e9
		}
		for i, p := range a {
			l := sort.SearchInts(ls, p.l) + 1
			if !(i < n-1 && a[i+1].r == p.r && a[i+1].l == p.l) {
				if superSize := t.pre(l); superSize < 2e9 {
					ans[p.i] = superSize - (p.r - p.l)
				}
			}
			t.update(l, p.l, p.r)
		}
		for _, v := range ans {
			Fprintln(out, v)
		}
	}
}

//func main() { cf2042D(bufio.NewReader(os.Stdin), os.Stdout) }
