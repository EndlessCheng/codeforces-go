package main

import (
	"bufio"
	. "fmt"
	"io"
)

var x915E = uint(1)

func fastRand915E() uint {
	x915E ^= x915E << 13
	x915E ^= x915E >> 17
	x915E ^= x915E << 5
	return x915E
}

type node915E struct {
	lr       [2]*node915E
	priority uint
	l, r     int
	work     bool
}

func (o *node915E) rotate(d int) *node915E {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap915E struct {
	root       *node915E
	workDayCnt int
	comparator func(a, b int) int
}

func newTreap915E() *treap915E {
	return &treap915E{comparator: func(a, b int) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *treap915E) _put(o *node915E, l, r int, work bool) *node915E {
	if o == nil {
		return &node915E{priority: fastRand915E(), l: l, r: r, work: work}
	}
	if cmp := t.comparator(l, o.l); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], l, r, work)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.work = work
	}
	return o
}

func (t *treap915E) put(l, r int, work bool) { t.root = t._put(t.root, l, r, work) }

func (t *treap915E) _delete(o *node915E, l int) *node915E {
	if o == nil {
		return nil
	}
	if cmp := t.comparator(l, o.l); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], l)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		cmp2 := 0
		if o.lr[0].priority > o.lr[1].priority {
			cmp2 = 1
		}
		o = o.rotate(cmp2)
		o.lr[cmp2] = t._delete(o.lr[cmp2], l)
	}
	return o
}

func (t *treap915E) delete(l int) { t.root = t._delete(t.root, l) }

func (t *treap915E) floor(l int) (floor *node915E) {
	for o := t.root; o != nil; {
		switch cmp := t.comparator(l, o.l); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *treap915E) next(l int) (next *node915E) {
	for o := t.root; o != nil; {
		if cmp := t.comparator(l, o.l); cmp != 0 {
			o = o.lr[1]
		} else {
			next = o
			o = o.lr[0]
		}
	}
	return
}

func (t *treap915E) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, work := o.r, o.work
		o.r = mid - 1
		t.put(mid, r, work)
	}
}

func (t *treap915E) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap915E) updateCnt(o *node915E, work bool) {
	if !o.work && work {
		t.workDayCnt += o.r - o.l + 1
	} else if o.work && !work {
		t.workDayCnt -= o.r - o.l + 1
	}
}

func (t *treap915E) merge(l, r int, work bool) {
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
func Sol915E(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	t := newTreap915E()
	t.put(1, n, true)
	t.workDayCnt = n
	for q := read(); q > 0; q-- {
		l, r, k := read(), read(), read()
		t.prepare(l, r)
		t.merge(l, r, k == 2)
		Fprintln(out, t.workDayCnt)
	}
}

//func main() {
//	Sol915E(os.Stdin, os.Stdout)
//}
