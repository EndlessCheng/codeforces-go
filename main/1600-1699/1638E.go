package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"time"
)

// https://github.com/EndlessCheng
type nodeM[K comparable, V any] struct {
	son      [2]*nodeM[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeM[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeM[K, V]) rotate(d int) *nodeM[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapM[K comparable, V any] struct {
	rd         uint
	root       *nodeM[K, V]
	comparator func(a, b K) int
}

func (t *treapM[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapM[K, V]) size() int   { return t.root.size() }
func (t *treapM[K, V]) empty() bool { return t.size() == 0 }

func (t *treapM[K, V]) _put(o *nodeM[K, V], key K, value V) *nodeM[K, V] {
	if o == nil {
		o = &nodeM[K, V]{priority: t.fastRand(), key: key, value: value}
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

func (t *treapM[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treapM[K, V]) _delete(o *nodeM[K, V], key K) *nodeM[K, V] {
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

func (t *treapM[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapM[K, V]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapM[K, V]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

func (t *treapM[K, V]) kth(k int) (o *nodeM[K, V]) {
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

func (t *treapM[K, V]) prev(key K) *nodeM[K, V]  { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapM[K, V]) next(key K) *nodeM[K, V]  { return t.kth(t.upperBoundIndex(key)) }
func (t *treapM[K, V]) floor(key K) *nodeM[K, V] { return t.kth(t.upperBoundIndex(key) - 1) }

func newMap[K cmp.Ordered, V any]() *treapM[K, V] {
	return &treapM[K, V]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

type fenwick []int

func (t fenwick) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += t[i]
	}
	return
}

func cf1638E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, c, x int
	var op string
	Fscan(in, &n, &m)

	type pair struct{ r, c int }
	t := newMap[int, pair]()
	t.put(1, pair{n, 1})
	t.put(n+1, pair{})
	split := func(mid int) {
		o := t.floor(mid)
		if o.key < mid {
			t.put(mid, o.value)
			o.value.r = mid - 1
		}
	}

	f := make(fenwick, n+1)
	inc := make([]int, n+1)

	for range m {
		Fscan(in, &op)
		if op[0] == 'C' {
			Fscan(in, &l, &r, &c)
			split(l)
			split(r + 1)
			for o := t.floor(l); o.key <= r; o = t.next(o.key) {
				d := inc[o.value.c] - inc[c]
				f.update(o.key, d)
				f.update(o.value.r+1, -d)
				t.delete(o.key)
			}
			t.put(l, pair{r, c})
		} else if op[0] == 'A' {
			Fscan(in, &c, &x)
			inc[c] += x
		} else {
			Fscan(in, &x)
			Fprintln(out, f.pre(x)+inc[t.floor(x).value.c])
		}
	}
}

//func main() { cf1638E(bufio.NewReader(os.Stdin), os.Stdout) }
