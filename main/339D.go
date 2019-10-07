package main

import (
	"bufio"
	. "fmt"
	"io"
)

type node339D struct {
	l, r   int
	val    int
	needOR bool
}
type segmentTree339D []node339D

func (t segmentTree339D) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	if lo.needOR {
		t[o].val = lo.val | ro.val
	} else {
		t[o].val = lo.val ^ ro.val
	}
}

func (t segmentTree339D) _build(arr []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = arr[l]
		t[o].needOR = true
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)
	t[o].needOR = !t[o<<1].needOR
	t._pushUp(o)
}

func (t segmentTree339D) _update(o, idx int, val int) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t segmentTree339D) init(arr []int)          { t._build(arr, 1, 1, len(arr)-1) }
func (t segmentTree339D) update(idx int, val int) { t._update(1, idx, val) }

func Sol339D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	n = 1 << uint(n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &arr[i])
	}

	t := make(segmentTree339D, 2*n)
	t.init(arr)
	for ; m > 0; m-- {
		var idx, val int
		Fscan(in, &idx, &val)
		t.update(idx, val)
		Fprintln(out, t[1].val)
	}
}

//func main() {
//	Sol339D(os.Stdin, os.Stdout)
//}
