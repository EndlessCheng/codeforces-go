package main

import (
	"bufio"
	. "fmt"
	"io"
)

type color bool

const red, black color = true, false

type rbNode struct {
	lr    [2]*rbNode
	msz   int
	key   int
	value int
	c     color
}

func (o *rbNode) isRed() bool {
	if o == nil {
		return false
	}
	return o.c == red
}

func (o *rbNode) pushUp() {
	msz := o.value
	if ol := o.lr[0]; ol != nil {
		msz += ol.msz
	}
	if or := o.lr[1]; or != nil {
		msz += or.msz
	}
	o.msz = msz
}

func (o *rbNode) rotate(d int) *rbNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	x.c = x.lr[d].c
	x.lr[d].c = red
	x.msz = o.msz
	o.pushUp()
	return x
}

func (o *rbNode) flipColors() {
	o.c = !o.c
	o.lr[0].c = !o.lr[0].c
	o.lr[1].c = !o.lr[1].c
}

func (o *rbNode) moveRedLeft() *rbNode {
	o.flipColors()
	if o.lr[1].lr[0].isRed() {
		o.lr[1] = o.lr[1].rotate(1)
		o = o.rotate(0)
		o.flipColors()
	}
	return o
}

func (o *rbNode) moveRedRight() *rbNode {
	o.flipColors()
	if o.lr[0].lr[0].isRed() {
		o = o.rotate(1)
		o.flipColors()
	}
	return o
}

func (o *rbNode) balance() *rbNode {
	if o.lr[1].isRed() {
		o = o.rotate(0)
	}
	if o.lr[0].isRed() && o.lr[0].lr[0].isRed() {
		o = o.rotate(1)
	}
	if o.lr[0].isRed() && o.lr[1].isRed() {
		o.flipColors()
	}
	o.pushUp()
	return o
}

func (o *rbNode) min() *rbNode {
	for o.lr[0] != nil {
		o = o.lr[0]
	}
	return o
}

func (o *rbNode) deleteMin() *rbNode {
	if o.lr[0] == nil {
		return nil
	}
	if !o.lr[0].isRed() && !o.lr[0].lr[0].isRed() {
		o = o.moveRedLeft()
	}
	o.lr[0] = o.lr[0].deleteMin()
	return o.balance()
}

type rbTree struct {
	root       *rbNode
	comparator func(a, b int) int
}

func newRBTree() *rbTree {
	return &rbTree{comparator: func(a, b int) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *rbTree) _put(o *rbNode, key int) *rbNode {
	if o == nil {
		return &rbNode{msz: 1, key: key, value: 1, c: red}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[1].isRed() && !o.lr[0].isRed() {
			o = o.rotate(0)
		}
		if o.lr[0].isRed() && o.lr[0].lr[0].isRed() {
			o = o.rotate(1)
		}
		if o.lr[0].isRed() && o.lr[1].isRed() {
			o.flipColors()
		}
	} else {
		o.value++
	}
	o.pushUp()
	return o
}

func (t *rbTree) put(key int) {
	t.root = t._put(t.root, key)
	t.root.c = black
}

func (t *rbTree) _delete(o *rbNode, key int) *rbNode {
	if cmp := t.comparator(key, o.key); cmp == 0 {
		if !o.lr[0].isRed() && !o.lr[0].lr[0].isRed() {
			o = o.moveRedLeft()
		}
		o.lr[0] = t._delete(o.lr[0], key)
	} else {
		if o.lr[0].isRed() {
			o = o.rotate(1)
		}
		if t.comparator(key, o.key) == -1 && o.lr[1] == nil {
			return nil
		}
		if !o.lr[1].isRed() && !o.lr[1].lr[0].isRed() {
			o = o.moveRedRight()
		}
		if t.comparator(key, o.key) == -1 {
			x := o.lr[1].min()
			o.key = x.key
			o.value = x.value
			o.lr[1] = o.lr[1].deleteMin()
		} else {
			o.lr[1] = t._delete(o.lr[1], key)
		}
	}
	return o.balance()
}

func (t *rbTree) get(key int) *rbNode {
	for o := t.root; o != nil; {
		if cmp := t.comparator(key, o.key); cmp >= 0 {
			o = o.lr[cmp]
		} else {
			return o
		}
	}
	return nil
}

func (t *rbTree) delete(key int) {
	if o := t.get(key); o != nil && o.value > 1 {
		o.value--
		return
	}
	if !t.root.lr[0].isRed() && !t.root.lr[1].isRed() {
		t.root.c = red
	}
	t.root = t._delete(t.root, key)
	if t.root != nil {
		t.root.c = black
	}
}

func (t *rbTree) floor(key int) (floor *rbNode) {
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

func (t *rbTree) next(key int) (next *rbNode) {
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

func (t *rbTree) max() (max *rbNode) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol527CRBTree(reader io.Reader, writer io.Writer) {
	cut := func(t, mt *rbTree, mid int) {
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

	w, h, mw, mh := newRBTree(), newRBTree(), newRBTree(), newRBTree()
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
//	Sol527CRBTree(os.Stdin, os.Stdout)
//}
