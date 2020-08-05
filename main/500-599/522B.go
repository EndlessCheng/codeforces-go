package main

import (
	"bufio"
	. "fmt"
	"io"
)

func max522B(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stNode522B struct {
	l, r int
	val  int
}
type segmentTree522B []stNode522B

func (t segmentTree522B) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = max522B(lo.val, ro.val)
}

func (t segmentTree522B) _build(arr []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = arr[l-1]
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)
	t._pushUp(o)
}

func (t segmentTree522B) _query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	mid := (t[o].l + t[o].r) >> 1
	res = -1e9
	if l <= mid {
		res = max522B(res, t._query(o<<1, l, r))
	}
	if mid < r {
		res = max522B(res, t._query(o<<1|1, l, r))
	}
	return
}

func (t segmentTree522B) init(arr []int) { t._build(arr, 1, 1, len(arr)) }
func (t segmentTree522B) query(l, r int) int {
	if l > r {
		return 0
	}
	return t._query(1, l, r)
}

// github.com/EndlessCheng/codeforces-go
func Sol522B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	sumW := 0
	w := make([]int, n)
	h := make([]int, n)
	for i := range w {
		Fscan(in, &w[i], &h[i])
		sumW += w[i]
	}

	t := make(segmentTree522B, 4*n)
	t.init(h)
	for i, wi := range w {
		maxH := max522B(t.query(1, i), t.query(i+2, n))
		Fprint(out, (sumW-wi)*maxH, " ")
	}
}

//func main() {
//	Sol522B(os.Stdin, os.Stdout)
//}
