package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type node00[K comparable] struct {
	son      [2]*node00[K]
	priority uint
	key      K
	subSize  int
}

func (o *node00[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node00[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *node00[K]) rotate(d int) *node00[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap00[K comparable] struct {
	rd         uint
	root       *node00[K]
	comparator func(a, b K) int
}

func (t *treap00[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap00[K]) size() int   { return t.root.size() }
func (t *treap00[K]) empty() bool { return t.size() == 0 }

func (t *treap00[K]) _put(o *node00[K], key K) *node00[K] {
	if o == nil {
		o = &node00[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c != 0 {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treap00[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treap00[K]) _delete(o *node00[K], key K) *node00[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.son[1] == nil {
			return o.son[0]
		}
		if o.son[0] == nil {
			return o.son[1]
		}
		d := 0
		if o.son[0].priority > o.son[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.son[d] = t._delete(o.son[d], key)
	}
	o.maintain()
	return o
}

func (t *treap00[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treap00[K]) min() *node00[K] { return t.kth(0) }

func (t *treap00[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treap00[K]) kth(k int) (o *node00[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func newSet00[K cmp.Ordered]() *treap00[K] {
	return &treap00[K]{
		rd:         1,
		comparator: cmp.Compare[K],
	}
}

type nodeM00[K comparable, V any] struct {
	son      [2]*nodeM00[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM00[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeM00[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeM00[K, V]) rotate(d int) *nodeM00[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapM00[K comparable, V any] struct {
	rd         uint
	root       *nodeM00[K, V]
	comparator func(a, b K) int
}

func (t *treapM00[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapM00[K, V]) size() int   { return t.root.size() }
func (t *treapM00[K, V]) empty() bool { return t.size() == 0 }

func (t *treapM00[K, V]) _put(o *nodeM00[K, V], key K, value V) *nodeM00[K, V] {
	if o == nil {
		o = &nodeM00[K, V]{priority: t.fastRand(), key: key, value: value}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.value = value
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key, value)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treapM00[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treapM00[K, V]) _delete(o *nodeM00[K, V], key K) *nodeM00[K, V] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.son[1] == nil {
			return o.son[0]
		}
		if o.son[0] == nil {
			return o.son[1]
		}
		d := 0
		if o.son[0].priority > o.son[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.son[d] = t._delete(o.son[d], key)
	}
	o.maintain()
	return o
}

func (t *treapM00[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapM00[K, V]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapM00[K, V]) kth(k int) (o *nodeM00[K, V]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func newMap00[K cmp.Ordered, V any]() *treapM00[K, V] {
	return &treapM00[K, V]{
		rd:         1,
		comparator: cmp.Compare[K],
	}
}

func cf2000H_2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v, m int
	var op string
	for Fscan(in, &T); T > 0; T-- {
		diff := newMap00[int, *treap00[int]]()
		put := func(l, r int) {
			k := r - l - 1
			o := diff.kth(diff.lowerBoundIndex(k))
			var t *treap00[int]
			if o == nil || o.key != k {
				t = newSet00[int]()
				diff.put(k, t)
			} else {
				t = o.value
			}
			t.put(l + 1)
		}
		del := func(l, r int) {
			k := r - l - 1
			t := diff.kth(diff.lowerBoundIndex(k)).value
			t.delete(l + 1)
			if t.empty() {
				diff.delete(k)
			}
		}

		set := newSet00[int]()
		set.put(0)
		pre := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			set.put(v)
			put(pre, v)
			pre = v
		}
		set.put(1e9)
		put(pre, 1e9)

		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &op, &v)
			if op == "+" {
				i := set.lowerBoundIndex(v)
				l, r := set.kth(i-1).key, set.kth(i).key
				set.put(v)
				del(l, r)
				put(l, v)
				put(v, r)
			} else if op == "-" {
				i := set.lowerBoundIndex(v)
				l, r := set.kth(i-1).key, set.kth(i+1).key
				set.delete(v)
				del(l, v)
				del(v, r)
				put(l, r)
			} else {
				Fprint(out, diff.kth(diff.lowerBoundIndex(v)).value.min().key, " ")
			}
		}
		Fprintln(out)
	}
}
