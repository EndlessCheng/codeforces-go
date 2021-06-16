package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"time"
)

// github.com/EndlessCheng/codeforces-go
type node98 struct {
	lr       [2]*node98
	priority uint
	key, sz  int
	s        int64
}

func (o *node98) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node98) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node98) sum() int64 {
	if o != nil {
		return o.s
	}
	return 0
}

func (o *node98) maintain() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
	o.s = int64(o.key) + o.lr[0].sum() + o.lr[1].sum()
}

func (o *node98) rotate(d int) *node98 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap98 struct {
	rd   uint
	root *node98
}

func (t *treap98) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap98) _put(o *node98, key int) *node98 {
	if o == nil {
		return &node98{priority: t.fastRand(), key: key, sz: 1, s: int64(key)}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	o.maintain()
	return o
}

func (t *treap98) put(key int) { t.root = t._put(t.root, key) }

func (t *treap98) _delete(o *node98, key int) *node98 {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
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
		o.lr[d] = t._delete(o.lr[d], key)
	}
	o.maintain()
	return o
}

func (t *treap98) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap98) kth(k int) (s int64) {
	for o := t.root; o != nil; {
		if ls := o.lr[0].size(); k < ls {
			o = o.lr[0]
		} else {
			s += o.lr[0].sum()
			if k > ls {
				s += int64(o.key)
			}
			k -= ls + 1
			if k < 0 {
				return
			}
			o = o.lr[1]
		}
	}
	return
}

type vi98 struct{ v, i int }
type mh98 []*vi98

func (h mh98) Len() int            { return len(h) }
func (h mh98) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h mh98) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *mh98) Push(v interface{}) { *h = append(*h, v.(*vi98)) }
func (h *mh98) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh98) push(v int) *vi98   { p := &vi98{v, len(*h)}; heap.Push(h, p); return p }

func CF1398E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := &treap98{rd: uint(time.Now().UnixNano())/2 + 1}
	h := mh98{}
	ptr := map[int]*vi98{}

	var q, tp, v int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &tp, &v)
		if v > 0 {
			t.put(v)
		} else {
			t.delete(-v)
		}
		if tp == 1 {
			if v > 0 {
				ptr[v] = h.push(v)
			} else {
				heap.Remove(&h, ptr[-v].i)
				delete(ptr, -v)
			}
		}
		if len(h) == 0 {
			Fprintln(out, t.root.sum())
		} else if len(h) == t.root.sz {
			Fprintln(out, t.root.sum()*2-int64(h[0].v))
		} else {
			mi := h[0].v
			t.delete(mi)
			extra := t.kth(t.root.sz - len(h))
			t.put(mi)
			Fprintln(out, t.root.sum()*2-extra-int64(mi))
		}
	}
}

//func main() { CF1398E(os.Stdin, os.Stdout) }
