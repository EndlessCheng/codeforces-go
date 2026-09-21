package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick364 []int

func (t fenwick364) reset(i int) {
	for ; i < len(t); i += i & -i {
		t[i] = 0
	}
}

func (t fenwick364) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], val)
	}
}

func (t fenwick364) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, t[i])
	}
	return
}

func p3364(in io.Reader, out io.Writer) {
	var n, lv, x, y, z int
	Fscan(in, &n)
	type data struct{ lv, str, int, atk, f int }
	d := make([]data, n)
	a := make([]int, 0, n*3)
	for i := range d {
		Fscan(in, &lv, &x, &y, &z)
		d[i] = data{lv, x, y, z, 1}
		a = append(a, x, y, z)
	}
	slices.SortFunc(d, func(a, b data) int { return a.lv - b.lv })

	slices.Sort(a)
	a = slices.Compact(a)
	for i, t := range d {
		d[i].str = sort.SearchInts(a, t.str)
		d[i].int = sort.SearchInts(a, t.int)
		d[i].atk = sort.SearchInts(a, t.atk)
	}

	// f[i] = max(f[j]) + 1
	// j < i
	// atk[j] <= str[i]
	// int[j] <= atk[i]

	t := make(fenwick364, len(a)+1)
	ans := 1
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		solve(l, mid)

		slices.SortFunc(d[l:mid], func(a, b data) int { return a.atk - b.atk })
		slices.SortFunc(d[mid:r], func(a, b data) int { return a.str - b.str })
		j := l
		for i := mid; i < r; i++ {
			for ; j < mid && d[j].atk <= d[i].str; j++ {
				t.update(d[j].int+1, d[j].f)
			}
			d[i].f = max(d[i].f, t.pre(d[i].atk+1)+1)
			ans = max(ans, d[i].f)
		}
		for j--; j >= l; j-- {
			t.reset(d[j].int + 1)
		}
		slices.SortFunc(d[l:r], func(a, b data) int { return a.lv - b.lv })

		solve(mid, r)
	}
	solve(0, n)
	Fprint(out, ans)
}

//func main() { p3364(bufio.NewReader(os.Stdin), os.Stdout) }
