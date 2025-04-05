package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"time"
)

// https://github.com/EndlessCheng
type nodeM70[K comparable, V any] struct {
	son      [2]*nodeM70[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM70[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeM70[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeM70[K, V]) rotate(d int) *nodeM70[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapM70[K comparable, V any] struct {
	rd         uint
	root       *nodeM70[K, V]
	comparator func(a, b K) int
}

func (t *treapM70[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapM70[K, V]) size() int   { return t.root.size() }
func (t *treapM70[K, V]) empty() bool { return t.size() == 0 }

func (t *treapM70[K, V]) _put(o *nodeM70[K, V], key K, value V) *nodeM70[K, V] {
	if o == nil {
		o = &nodeM70[K, V]{priority: t.fastRand(), key: key, value: value}
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

func (t *treapM70[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treapM70[K, V]) _delete(o *nodeM70[K, V], key K) *nodeM70[K, V] {
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

func (t *treapM70[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapM70[K, V]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapM70[K, V]) kth(k int) (o *nodeM70[K, V]) {
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

func newMap70[K cmp.Ordered, V any]() *treapM70[K, V] {
	return &treapM70[K, V]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[K],
	}
}

func cf70D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	det := func(x1, y1, x2, y2 int) int { return x1*y2 - x2*y1 }
	_remove := func(t *treapM70[int, int], i int) bool {
		if i <= 0 || i+1 >= t.size() {
			return false
		}
		pre := t.kth(i - 1)
		cur := t.kth(i)
		nxt := t.kth(i + 1)
		if det(cur.key-pre.key, cur.value-pre.value, nxt.key-pre.key, nxt.value-pre.value) >= 0 {
			t.delete(cur.key)
			return true
		}
		return false
	}
	contains := func(t *treapM70[int, int], x, y int) bool {
		i := t.lowerBoundIndex(x)
		if i == t.size() {
			return false
		}
		cur := t.kth(i)
		if cur.key == x {
			return y <= cur.value
		}
		if i == 0 {
			return false
		}
		pre := t.kth(i - 1)
		return det(cur.key-pre.key, cur.value-pre.value, x-pre.key, y-pre.value) <= 0
	}
	insert := func(t *treapM70[int, int], x, y int) bool {
		if contains(t, x, y) {
			return false
		}
		t.put(x, y)
		idx := t.lowerBoundIndex(x)
		for j := idx + 1; _remove(t, j); {
		}
		for j := idx - 1; _remove(t, j); j-- {
		}
		return true
	}
	top := newMap70[int, int]()  // 上凸包
	down := newMap70[int, int]() // 下凸包（y 取反）

	var q, op, x, y int
	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &x, &y)
		if op == 1 {
			insert(top, x, y)
			insert(down, x, -y)
		} else if contains(top, x, y) && contains(down, x, -y) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf70D(bufio.NewReader(os.Stdin), os.Stdout) }
