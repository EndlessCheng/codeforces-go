package main

import (
	"bufio"
	. "fmt"
	"io"
)

type node struct {
	l, r        int
	sum         int
	revChildren bool
}
type segTree []node

func (t segTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = lo.sum + ro.sum
}

func (t segTree) _build(arr []int, pos uint8, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = arr[l-1] >> pos & 1
		return
	}
	mid := (l + r) >> 1
	t._build(arr, pos, o<<1, l, mid)
	t._build(arr, pos, o<<1|1, mid+1, r)
	t._pushUp(o)
}

func (t segTree) _spread(o int) {
	if t[o].revChildren {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum = lo.r - lo.l + 1 - lo.sum
		ro.sum = ro.r - ro.l + 1 - ro.sum
		lo.revChildren = !lo.revChildren
		ro.revChildren = !ro.revChildren
		t[o].revChildren = false
	}
}

func (t segTree) _rev(o, l, r int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum = or - ol + 1 - t[o].sum
		t[o].revChildren = !t[o].revChildren
		return
	}
	t._spread(o)
	mid := (ol + or) >> 1
	if l <= mid {
		t._rev(o<<1, l, r)
	}
	if mid < r {
		t._rev(o<<1|1, l, r)
	}
	t._pushUp(o)
}

func (t segTree) _query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t._spread(o)
	mid := (t[o].l + t[o].r) >> 1
	if l <= mid {
		res += t._query(o<<1, l, r)
	}
	if mid < r {
		res += t._query(o<<1|1, l, r)
	}
	return
}

func (t segTree) init(arr []int, pos uint8) { t._build(arr, pos, 1, 1, len(arr)) }
func (t segTree) rev(l, r int)              { t._rev(1, l, r) }
func (t segTree) query(l, r int) int        { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func Sol242E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q int
	Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	trees := make([]segTree, 20)
	for i := range trees {
		trees[i] = make(segTree, 4*n)
		trees[i].init(arr, uint8(i))
	}

	for Fscan(in, &q); q > 0; q-- {
		var op, l, r, x int
		Fscan(in, &op, &l, &r)
		if op == 1 {
			sum := int64(0)
			for i, t := range trees {
				sum += 1 << uint(i) * int64(t.query(l, r))
			}
			Fprintln(out, sum)
		} else {
			Fscan(in, &x)
			for i, t := range trees {
				if x>>uint(i)&1 == 1 {
					t.rev(l, r)
				}
			}
		}
	}
}

//func main() {
//	Sol242E(os.Stdin, os.Stdout)
//}
