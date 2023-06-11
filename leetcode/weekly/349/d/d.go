package main

import (
	"sort"
	"time"
)

// https://space.bilibili.com/206214
type node struct {
	lr       [2]*node
	priority uint
	key      int
	value    int
	subCnt   int
	mx       int
}

func (o *node) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node) MX() int {
	if o != nil {
		return o.mx
	}
	return 0
}

func (o *node) maintain() {
	o.subCnt = 1 + o.lr[0].size() + o.lr[1].size()
	o.mx = max(o.value, max(o.lr[0].MX(), o.lr[1].MX()))
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap struct {
	rd   uint
	root *node
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) size() int { return t.root.size() }

func (t *treap) _put(o *node, key int, value int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, value: value, subCnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, value)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.value = value
	}
	o.maintain()
	return o
}

func (t *treap) put(key int, value int) { t.root = t._put(t.root, key, value) }

func (t *treap) _delete(o *node, key int) *node {
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

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap) get(key int) *node {
	for o := t.root; o != nil; {
		if c := o.cmp(key); c >= 0 {
			o = o.lr[c]
		} else {
			return o
		}
	}
	return nil // 
}

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }

func (t *treap) ge(key int) (mx int) {
	mx = -1
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			mx = max(mx, max(o.value, o.lr[1].MX()))
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			mx = max(mx, max(o.value, o.lr[1].MX()))
			return
		}
	}
	return
}

func (o *node) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

type stNode struct {
	lo, ro   *stNode
	l, r, mx int
}

func (o *stNode) get() int {
	if o != nil {
		return o.mx
	}
	return -1
}

func (o *stNode) update(i, val int) {
	if o.l == o.r {
		o.mx = max(o.mx, val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &stNode{l: o.l, r: m}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == nil {
			o.ro = &stNode{l: m + 1, r: o.r}
		}
		o.ro.update(i, val)
	}
	o.mx = max(o.lo.get(), o.ro.get())
}

func (o *stNode) query(l, r int) int {
	if o == nil || l > o.r || r < o.l {
		return -1
	}
	if l <= o.l && o.r <= r {
		return o.mx
	}
	return max(o.lo.query(l, r), o.ro.query(l, r))
}

func maximumSumQueries2(nums1, nums2 []int, queries [][]int) (ans []int) {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, v := range nums1 {
		a[i] = pair{v, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][0] > queries[j][0] })

	ans = make([]int, len(queries))
	rt := &stNode{l: 0, r: 2e9}
	i := len(a) - 1
	for _, q := range queries {
		for i >= 0 && a[i].x >= q[0] {
			rt.update(a[i].y, a[i].x+a[i].y)
			i--
		}
		ans[q[2]] = rt.query(q[1], 2e9)
	}
	return
}

// https://space.bilibili.com/206214
func maximumSumQueries(nums1, nums2 []int, queries [][]int) (ans []int) {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, v := range nums1 {
		a[i] = pair{v, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][0] > queries[j][0] })

	ans = make([]int, len(queries))
	st := []pair{}
	i := len(a) - 1
	for _, q := range queries {
		for i >= 0 && a[i].x >= q[0] {
			for len(st) > 0 && st[len(st)-1].y <= a[i].x+a[i].y {
				st = st[:len(st)-1]
			}
			if len(st) == 0 || st[len(st)-1].x < a[i].y {
				st = append(st, pair{a[i].y, a[i].x + a[i].y})
			}
			i--
		}
		j := sort.Search(len(st), func(i int) bool { return st[i].x >= q[1] })
		if j < len(st) {
			ans[q[2]] = st[j].y
		} else {
			ans[q[2]] = -1
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
