package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"math/bits"
	"time"
)

// https://space.bilibili.com/206214
type nodeM320[K comparable, V any] struct {
	son      [2]*nodeM320[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM320[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeM320[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeM320[K, V]) rotate(d int) *nodeM320[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapM320[K comparable, V any] struct {
	rd         uint
	root       *nodeM320[K, V]
	comparator func(a, b K) int
}

func (t *treapM320[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapM320[K, V]) size() int   { return t.root.size() }
func (t *treapM320[K, V]) empty() bool { return t.size() == 0 }

func (t *treapM320[K, V]) _put(o *nodeM320[K, V], key K, value V) *nodeM320[K, V] {
	if o == nil {
		o = &nodeM320[K, V]{priority: t.fastRand(), key: key, value: value}
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

func (t *treapM320[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treapM320[K, V]) _delete(o *nodeM320[K, V], key K) *nodeM320[K, V] {
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

func (t *treapM320[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapM320[K, V]) min() *nodeM320[K, V] { return t.kth(0) }
func (t *treapM320[K, V]) max() *nodeM320[K, V] { return t.kth(t.size() - 1) }

func (t *treapM320[K, V]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapM320[K, V]) upperBoundIndex(key K) (kth int) {
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

func (t *treapM320[K, V]) kth(k int) (o *nodeM320[K, V]) {
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

func (t *treapM320[K, V]) prev(key K) *nodeM320[K, V] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapM320[K, V]) next(key K) *nodeM320[K, V] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapM320[K, V]) find(key K) *nodeM320[K, V] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newMap320[K cmp.Ordered, V any]() *treapM320[K, V] {
	return &treapM320[K, V]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[K],
	}
}

func p3320(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, dfn int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		var v, w, wt int
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
	vToDFN := make([]int, n)
	var build func(int, int)
	build = func(v, p int) {
		vToDFN[v] = dfn
		dfn++
		pa[v][0] = p
		for _, e := range g[v] {
			w := e.to
			if w != p {
				dep[w] = dep[v] + 1
				dis[w] = dis[v] + e.wt
				build(w, v)
			}
		}
	}
	build(0, -1)
	for i := 0; i < mx-1; i++ {
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

	ans := 0
	t := newMap320[int, int]()
	has := make([]bool, n)
	for ; m > 0; m-- {
		var v int
		Fscan(in, &v)
		v--
		d := vToDFN[v]
		if has[d] {
			t.delete(d)
		} else {
			t.put(d, v)
		}
		has[d] = !has[d]
		if t.size() < 2 {
			Fprintln(out, 0)
			continue
		}

		o := t.prev(d)
		if o == nil {
			o = t.max()
		}
		p := o.value
		o = t.next(d)
		if o == nil {
			o = t.min()
		}
		q := o.value
		inc := getDis(p, v) + getDis(v, q) - getDis(p, q)
		if has[d] {
			ans += inc
		} else {
			ans -= inc
		}
		Fprintln(out, ans)
	}
}

//func main() { p3320(bufio.NewReader(os.Stdin), os.Stdout) }
