package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type seg74 []struct{ l, r, mx, p int }

func (seg74) maxPos(a, pa, b, pb int) (int, int) {
	if a > b {
		return a, pa
	}
	return b, pb
}

func (t seg74) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg74) update(o, i, v, p int) {
	if t[o].l == t[o].r {
		t[o].mx, t[o].p = v, p
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, v, p)
	} else {
		t.update(o<<1|1, i, v, p)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].mx, t[o].p = t.maxPos(lo.mx, lo.p, ro.mx, ro.p)
}

func (t seg74) query(o, l, r int) (int, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx, t[o].p
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	mxl, pl := t.query(o<<1, l, r)
	mxr, pr := t.query(o<<1|1, l, r)
	return t.maxPos(mxl, pl, mxr, pr)
}

func CF474E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, mxL, end int
	var d int64
	Fscan(in, &n, &d)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := append([]int64(nil), a...)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	k := 1
	kth := map[int64]int{b[0]: k}
	for i := 1; i < n; i++ {
		if b[i] != b[i-1] {
			k++
			kth[b[i]] = k
		}
	}

	t := make(seg74, 4*k)
	t.build(1, 1, k)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}
	for i, v := range a {
		cur := 0
		if j := sort.Search(n, func(i int) bool { return b[i] > v-d }); j > 0 {
			if mx, pre := t.query(1, 1, kth[b[j-1]]); mx > cur {
				cur = mx
				fa[i] = pre
			}
		}
		if j := sort.Search(n, func(i int) bool { return b[i] >= v+d }); j < n {
			if mx, pre := t.query(1, kth[b[j]], k); mx > cur {
				cur = mx
				fa[i] = pre
			}
		}
		cur++
		if cur > mxL {
			mxL, end = cur, i
		}
		t.update(1, kth[v], cur, i)
	}
	ans := []int{}
	for x := end; x >= 0; x = fa[x] {
		ans = append(ans, x)
	}
	Fprintln(out, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i]+1, " ")
	}
}

//func main() { CF474E(os.Stdin, os.Stdout) }
