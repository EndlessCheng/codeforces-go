package main

import (
	"bufio"
	. "fmt"
	"io"
)

type node42 struct {
	l, r        int32
	sum         int32
	revChildren bool
}
type seg42 []node42

func (t seg42) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = lo.sum + ro.sum
}

func (t seg42) _build(arr []int32, pos uint8, o int, l, r int32) {
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

func (t seg42) _spread(o int) {
	if t[o].revChildren {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum = lo.r - lo.l + 1 - lo.sum
		ro.sum = ro.r - ro.l + 1 - ro.sum
		lo.revChildren = !lo.revChildren
		ro.revChildren = !ro.revChildren
		t[o].revChildren = false
	}
}

func (t seg42) _rev(o int, l, r int32) {
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

func (t seg42) _query(o int, l, r int32) (res int32) {
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

func (t seg42) init(arr []int32, pos uint8) { t._build(arr, pos, 1, 1, int32(len(arr))) }
func (t seg42) rev(l, r int32)              { t._rev(1, l, r) }
func (t seg42) query(l, r int32) int32      { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func Sol242E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q int
	Fscan(in, &n)
	arr := make([]int32, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	trees := make([]seg42, 20)
	for i := range trees {
		trees[i] = make(seg42, 4*n)
		trees[i].init(arr, uint8(i))
	}

	for Fscan(in, &q); q > 0; q-- {
		var op, l, r, x int32
		Fscan(in, &op, &l, &r)
		if op == 1 {
			sum := 0
			for i, t := range trees {
				sum += 1 << uint(i) * int(t.query(l, r))
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
