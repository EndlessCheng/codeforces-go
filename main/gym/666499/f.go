package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"os"
	"slices"
	"time"
)

// https://github.com/EndlessCheng
type node[K comparable] struct {
	son      [2]*node[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
	keySum   int
	subSum   int
}

func (o *node[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node[K]) getSum() int {
	if o != nil {
		return o.subSum
	}
	return 0
}

func (o *node[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
	o.subSum = o.keySum + o.son[0].getSum() + o.son[1].getSum()
}

func (o *node[K]) rotate(d int) *node[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap[K comparable] struct {
	rd         uint
	root       *node[K]
	comparator func(a, b K) int
	keyToInt   func(key K) int
}

func (t *treap[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap[K]) size() int   { return t.root.size() }
func (t *treap[K]) empty() bool { return t.size() == 0 }

func (t *treap[K]) _put(o *node[K], key K, num int) *node[K] {
	if o == nil {
		if num < 0 {
			panic(num)
		}
		o = &node[K]{priority: t.fastRand(), key: key, keyCnt: num, keySum: t.keyToInt(key) * num}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt += num
			if o.keyCnt < 0 {
				panic(o.keyCnt)
			}
			o.keySum += t.keyToInt(key) * num
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key, num)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treap[K]) put(key K, num int) { t.root = t._put(t.root, key, num) }

func (t *treap[K]) preSum(k int) (sum int) {
	for o := t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
			continue
		}
		sum += o.son[0].getSum()
		k -= leftSize
		if k <= o.keyCnt {
			sum += t.keyToInt(o.key) * k
			break
		}
		sum += o.keySum
		k -= o.keyCnt
		o = o.son[1]
	}
	return
}

func newTreap() *treap[int] {
	return &treap[int]{
		rd:         uint(time.Now().UnixNano()),
		comparator: func(a, b int) int { return b - a },
		keyToInt:   func(key int) int { return key },
	}
}

func cfF(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
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

		if n < 3 {
			slices.Sort(a)
			ans := 0
			for _, v := range a[n-k:] {
				ans += v
			}
			Fprintln(out, ans)
			continue
		}

		sum := make([]int, n)
		subRes := make([]struct{ maxD, leaf, maxD2, leaf2, bestSon int }, n)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			type info struct{ d, l, son int }
			b := []info{}
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v)
					b = append(b, info{subRes[w].maxD + 1, subRes[w].leaf, w})
				}
			}
			if len(b) == 0 {
				sum[v] = a[v]
				subRes[v].leaf = v
				return
			}
			slices.SortFunc(b, func(a, b info) int { return cmp.Or(b.d-a.d, a.l-b.l) })
			sum[b[0].l] += a[v]
			subRes[v].maxD = b[0].d
			subRes[v].leaf = b[0].l
			subRes[v].bestSon = b[0].son
			if len(b) > 1 {
				subRes[v].maxD2 = b[1].d
				subRes[v].leaf2 = b[1].l
			} else {
				subRes[v].leaf2 = -1
			}
		}

		rt := 0
		for i, to := range g {
			if len(to) > 1 {
				rt = i
				break
			}
		}
		dfs(rt, -1)

		t := newTreap()
		for _, v := range sum {
			t.put(v, 1)
		}

		ans := 0
		var reroot func(int, int, int, int)
		reroot = func(v, fa, faD, faLeaf int) {
			sub := subRes[v]
			if sub.maxD == 0 {
				t.put(a[v], -1)
				ans = max(ans, a[v]+t.preSum(k-1))
				t.put(a[v], 1)
				return
			}

			t.put(sum[sub.leaf], -1)
			sum[sub.leaf] -= a[v]
			t.put(sum[sub.leaf], 1)
			ans = max(ans, a[v]+t.preSum(k-1))

			for _, w := range g[v] {
				if w == fa {
					continue
				}

				subD, leaf := sub.maxD, sub.leaf
				if w == sub.bestSon {
					subD, leaf = sub.maxD2, sub.leaf2
				}
				if faD > subD || faD == subD && faLeaf < leaf {
					subD, leaf = faD, faLeaf
				}

				t.put(sum[leaf], -1)
				sum[leaf] += a[v]
				t.put(sum[leaf], 1)

				reroot(w, v, subD+1, leaf)

				t.put(sum[leaf], -1)
				sum[leaf] -= a[v]
				t.put(sum[leaf], 1)
			}

			t.put(sum[sub.leaf], -1)
			sum[sub.leaf] += a[v]
			t.put(sum[sub.leaf], 1)
		}
		reroot(rt, -1, -1, -1)
		Fprintln(out, ans)
	}
}

func main() { cfF(bufio.NewReader(os.Stdin), os.Stdout) }
