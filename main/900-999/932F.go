package main

import (
	. "fmt"
	"io"
	"math/big"
	"sort"
)

// https://github.com/EndlessCheng
type node32[K comparable, V any] struct {
	son      [2]*node32[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *node32[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node32[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *node32[K, V]) rotate(d int) *node32[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap32[K comparable, V any] struct {
	rd         uint
	root       *node32[K, V]
	comparator func(a, b K) int
}

func (t *treap32[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap32[K, V]) size() int   { return t.root.size() }
func (t *treap32[K, V]) empty() bool { return t.size() == 0 }

func (t *treap32[K, V]) _put(o *node32[K, V], key K, value V) *node32[K, V] {
	if o == nil {
		o = &node32[K, V]{priority: t.fastRand(), key: key, value: value}
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

func (t *treap32[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treap32[K, V]) _delete(o *node32[K, V], key K) *node32[K, V] {
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

func (t *treap32[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treap32[K, V]) lowerBoundIndex(key K) (kth int) {
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

func (t *treap32[K, V]) kth(k int) (o *node32[K, V]) {
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

type vec32 struct{ x, y int }

func (a vec32) dot(b vec32) int { return a.x*b.x + a.y*b.y }

func cf932F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	det := func(x1, y1, x2, y2 int) int {
		v := new(big.Int).Mul(big.NewInt(int64(x1)), big.NewInt(int64(y2)))
		w := new(big.Int).Mul(big.NewInt(int64(y1)), big.NewInt(int64(x2)))
		return v.Cmp(w)
	}
	remove := func(t *treap32[int, int], i int) bool {
		if i <= 0 || i+1 >= t.size() {
			return false
		}
		pre := t.kth(i - 1)
		cur := t.kth(i)
		nxt := t.kth(i + 1)
		if det(cur.key-pre.key, cur.value-pre.value, nxt.key-pre.key, nxt.value-pre.value) <= 0 {
			t.delete(cur.key)
			return true
		}
		return false
	}
	contains := func(t *treap32[int, int], x, y int) bool {
		i := t.lowerBoundIndex(x)
		if i == t.size() {
			return false
		}
		cur := t.kth(i)
		if cur.key == x {
			return y >= cur.value
		}
		if i == 0 {
			return false
		}
		pre := t.kth(i - 1)
		return det(cur.key-pre.key, cur.value-pre.value, x-pre.key, y-pre.value) >= 0
	}
	insert := func(t *treap32[int, int], x, y int) {
		if contains(t, x, y) {
			return
		}
		t.put(x, y)
		idx := t.lowerBoundIndex(x)
		for j := idx + 1; remove(t, j); {
		}
		for j := idx - 1; remove(t, j); j-- {
		}
	}

	ans := make([]any, n)
	var dfs func(int, int) *treap32[int, int]
	dfs = func(v, fa int) *treap32[int, int] {
		t := &treap32[int, int]{
			rd:         1,
			comparator: func(a, b int) int { return a - b },
		}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			tw := dfs(w, v)
			if t.size() < tw.size() {
				t, tw = tw, t
			}
			var f func(*node32[int, int])
			f = func(o *node32[int, int]) {
				if o == nil {
					return
				}
				f(o.son[0])
				insert(t, o.key, o.value)
				f(o.son[1])
			}
			f(tw.root)
		}
		f := 0
		if !t.empty() {
			p := vec32{a[v], 1}
			j := sort.Search(t.size()-1, func(j int) bool {
				q := t.kth(j)
				q2 := t.kth(j + 1)
				return p.dot(vec32{q.key, q.value}) < p.dot(vec32{q2.key, q2.value})
			})
			q := t.kth(j)
			f = p.dot(vec32{q.key, q.value})
		}
		ans[v] = f
		insert(t, b[v], f)
		return t
	}
	dfs(0, -1)
	Fprintln(out, ans...)
}

//func main() { cf932F(bufio.NewReader(os.Stdin), os.Stdout) }
