package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

var _x527C = uint(time.Now().UnixNano())

func fastRand527C() uint {
	_x527C ^= _x527C << 13
	_x527C ^= _x527C >> 17
	_x527C ^= _x527C << 5
	return _x527C
}

type node527C struct {
	lr       [2]*node527C
	priority uint
	msz      int
	key      int
	value    int
}

func (o *node527C) pushUp() {
	msz := o.value
	if ol := o.lr[0]; ol != nil {
		msz += ol.msz
	}
	if or := o.lr[1]; or != nil {
		msz += or.msz
	}
	o.msz = msz
}

type treap527C struct {
	root       *node527C
	comparator func(a, b int) int
}

func newTreap527C() *treap527C {
	return &treap527C{comparator: func(a, b int) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *treap527C) rotate(o *node527C, d int) *node527C {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	x.msz = o.msz
	o.pushUp()
	return x
}

func (t *treap527C) _put(o *node527C, key int) *node527C {
	if o == nil {
		return &node527C{priority: fastRand527C(), msz: 1, key: key, value: 1}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = t.rotate(o, cmp^1)
		}
	} else {
		o.value++
	}
	o.pushUp()
	return o
}

func (t *treap527C) put(key int) { t.root = t._put(t.root, key) }

func (t *treap527C) _delete(o *node527C, key int) *node527C {
	if o == nil {
		return nil
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		if o.value > 1 {
			o.value--
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
			o = t.rotate(o, cmp2)
			o.lr[cmp2] = t._delete(o.lr[cmp2], key)
		}
	}
	o.pushUp()
	return o
}

func (t *treap527C) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap527C) floor(key int) (floor *node527C) {
	for o := t.root; o != nil; {
		switch cmp := t.comparator(key, o.key); {
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

func (t *treap527C) next(key int) (next *node527C) {
	for o := t.root; o != nil; {
		if cmp := t.comparator(key, o.key); cmp != 0 {
			o = o.lr[1]
		} else {
			next = o
			o = o.lr[0]
		}
	}
	return
}

func (t *treap527C) max() (max *node527C) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol527CTreap(reader io.Reader, writer io.Writer) {
	cut := func(t, mt *treap527C, mid int) {
		o := t.floor(mid)
		l := o.key
		r := t.next(l).key
		mt.delete(r - l)
		mt.put(r - mid)
		mt.put(mid - l)
		t.put(mid)
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	w, h, mw, mh := newTreap527C(), newTreap527C(), newTreap527C(), newTreap527C()
	var w0, h0, n int
	Fscan(in, &w0, &h0, &n)
	w.put(0)
	w.put(w0)
	mw.put(w0)
	h.put(0)
	h.put(h0)
	mh.put(h0)
	for ; n > 0; n-- {
		var op string
		var mid int
		Fscan(in, &op, &mid)
		if op[0] == 'V' {
			cut(w, mw, mid)
		} else {
			cut(h, mh, mid)
		}
		Fprintln(out, int64(mw.max().key)*int64(mh.max().key))
	}
}

//func main() {
//	Sol527CTreap(os.Stdin, os.Stdout)
//}
