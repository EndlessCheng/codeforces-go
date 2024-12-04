package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
	"sort"
	"time"
)

// https://github.com/EndlessCheng
type node46[K comparable] struct {
	son      [2]*node46[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
}

func (o *node46[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node46[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
}

func (o *node46[K]) rotate(d int) *node46[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap46[K comparable] struct {
	rd         uint
	root       *node46[K]
	comparator func(a, b K) int
}

func (t *treap46[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap46[K]) size() int   { return t.root.size() }
func (t *treap46[K]) empty() bool { return t.size() == 0 }

func (t *treap46[K]) _put(o *node46[K], key K) *node46[K] {
	if o == nil {
		o = &node46[K]{priority: t.fastRand(), key: key, keyCnt: 1}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt++
		} else {
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

func (t *treap46[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treap46[K]) _delete(o *node46[K], key K) *node46[K] {
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
		if o.keyCnt > 1 {
			o.keyCnt--
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
	}
	o.maintain()
	return o
}

func (t *treap46[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treap46[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treap46[K]) at(k int) (o *node46[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + o.keyCnt
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func newTreap46[K cmp.Ordered]() *treap46[K] {
	return &treap46[K]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[K],
	}
}

func cf2046C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ys := map[int][]int{}
		for range n {
			var x, y int
			Fscan(in, &x, &y)
			ys[x] = append(ys[x], y)
		}

		suf := newTreap46[int]()
		xs := make([]int, 0, len(ys))
		for x, ys := range ys {
			xs = append(xs, x)
			for _, y := range ys {
				suf.put(y)
			}
		}
		slices.Sort(xs)

		ans, xx, yy := 0, int(-1e9), int(-1e9)
		pre := newTreap46[int]()
		for _, x := range xs {
			if !pre.empty() {
				sort.Search(n/4, func(low int) bool {
					low++
					o := pre.at(low - 1)
					if o == nil {
						return true
					}
					minY := o.key
					o = suf.at(low - 1)
					if o == nil {
						return true
					}
					minY = max(minY, o.key) + 1

					i := pre.lowerBoundIndex(minY)
					if pre.size()-i < low {
						return true
					}
					i = suf.lowerBoundIndex(minY)
					if suf.size()-i < low {
						return true
					}
				
					if low > ans {
						ans = low
						xx, yy = x, minY
					}
					return false
				})
			}
			for _, y := range ys[x] {
				pre.put(y)
				suf.delete(y)
			}
		}
		Fprintln(out, ans)
		Fprintln(out, xx, yy)
	}
}

//func main() { cf2046C(bufio.NewReader(os.Stdin), os.Stdout) }
