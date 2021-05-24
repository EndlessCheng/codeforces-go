package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

type node15 struct {
	lr       [2]*node15
	priority uint
	l, r     int
	work     bool
}

func (o *node15) cmp(b int) int {
	switch {
	case b < o.l:
		return 0
	case b > o.l:
		return 1
	default:
		return -1
	}
}

func (o *node15) rotate(d int) *node15 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap15 struct {
	rd         uint
	root       *node15
	workDayCnt int
}

func (t *treap15) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap15) _put(o *node15, l, r int, work bool) *node15 {
	if o == nil {
		return &node15{priority: t.fastRand(), l: l, r: r, work: work}
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._put(o.lr[d], l, r, work)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.work = work
	}
	return o
}

func (t *treap15) put(l, r int, work bool) { t.root = t._put(t.root, l, r, work) }

func (t *treap15) _delete(o *node15, l int) *node15 {
	if o == nil {
		return nil
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], l)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		d = 0
		if o.lr[0].priority > o.lr[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.lr[d] = t._delete(o.lr[d], l)
	}
	return o
}

func (t *treap15) delete(l int) { t.root = t._delete(t.root, l) }

func (t *treap15) floor(key int) (floor *node15) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *treap15) next(l int) (next *node15) {
	for o := t.root; o != nil; {
		if o.cmp(l) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func (t *treap15) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, work := o.r, o.work
		o.r = mid - 1
		t.put(mid, r, work)
	}
}

func (t *treap15) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap15) updateCnt(o *node15, work bool) {
	if !o.work && work {
		t.workDayCnt += o.r - o.l + 1
	} else if o.work && !work {
		t.workDayCnt -= o.r - o.l + 1
	}
}

func (t *treap15) merge(l, r int, work bool) {
	t.prepare(l, r)
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		t.updateCnt(o, work)
		t.delete(o.l)
	}
	o := t.floor(l)
	t.updateCnt(o, work)
	o.r = r
	o.work = work
}

// github.com/EndlessCheng/codeforces-go
func CF915E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, k int
	Fscan(in, &n, &q)
	t := &treap15{rd: uint(time.Now().UnixNano())/2 + 1}
	t.put(1, n, true)
	t.workDayCnt = n
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &k)
		t.merge(l, r, k == 2)
		Fprintln(out, t.workDayCnt)
	}
}

//func main() { CF915E(os.Stdin, os.Stdout) }
