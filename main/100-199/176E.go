package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"math/bits"
	"time"
)

// https://github.com/EndlessCheng
type nodeM76[K comparable, V any] struct {
	son      [2]*nodeM76[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM76[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeM76[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeM76[K, V]) rotate(d int) *nodeM76[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapM76[K comparable, V any] struct {
	rd         uint
	root       *nodeM76[K, V]
	comparator func(a, b K) int
}

func (t *treapM76[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapM76[K, V]) size() int   { return t.root.size() }
func (t *treapM76[K, V]) empty() bool { return t.size() == 0 }

func (t *treapM76[K, V]) _put(o *nodeM76[K, V], key K, value V) *nodeM76[K, V] {
	if o == nil {
		o = &nodeM76[K, V]{priority: t.fastRand(), key: key, value: value}
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

func (t *treapM76[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treapM76[K, V]) _delete(o *nodeM76[K, V], key K) *nodeM76[K, V] {
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

func (t *treapM76[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapM76[K, V]) min() *nodeM76[K, V] { return t.kth(0) }
func (t *treapM76[K, V]) max() *nodeM76[K, V] { return t.kth(t.size() - 1) }

func (t *treapM76[K, V]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapM76[K, V]) upperBoundIndex(key K) (kth int) {
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

func (t *treapM76[K, V]) kth(k int) (o *nodeM76[K, V]) {
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

func (t *treapM76[K, V]) prev(key K) *nodeM76[K, V] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapM76[K, V]) next(key K) *nodeM76[K, V] { return t.kth(t.upperBoundIndex(key)) }

func newMap76[K cmp.Ordered, V any]() *treapM76[K, V] {
	return &treapM76[K, V]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func cf176E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v, w, wt, ans, ts int
	var op string
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for range n - 1 {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dis := make([]int, n)
	dfn := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, p int) {
		ts++
		dfn[v] = ts
		pa[v][0] = p
		for _, e := range g[v] {
			w := e.to
			if w == p {
				continue
			}
			dep[w] = dep[v] + 1
			dis[w] = dis[v] + e.wt
			dfs(w, v)
		}
	}
	dfs(0, -1)
	for i := range mx - 1 {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	getDis := func(v, w int) int { return dis[v] + dis[w] - dis[getLCA(v, w)]*2 }

	t := newMap76[int, int]()
	do := func(d, mul int) {
		if t.empty() {
			return
		}
		o := t.prev(d)
		if o == nil {
			o = t.max()
		}
		pre := o.value
		o = t.next(d)
		if o == nil {
			o = t.min()
		}
		nxt := o.value
		ans += (getDis(pre, v) + getDis(v, nxt) - getDis(pre, nxt)) * mul
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &op)
		if op == "?" {
			Fprintln(out, ans/2)
			continue
		}
		Fscan(in, &v)
		v--
		d := dfn[v]
		if op == "+" {
			do(d, 1)
			t.put(d, v)
		} else {
			t.delete(d)
			do(d, -1)
		}
	}
}

//func main() { cf176E(bufio.NewReader(os.Stdin), os.Stdout) }
