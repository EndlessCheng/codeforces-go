package main

import (
	"sort"
	"time"
)

type node struct {
	lr       [2]*node
	priority uint
	key      int
	dupCnt   int
	sz       int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() { o.sz = o.dupCnt + o.lr[0].size() + o.lr[1].size() }

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

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, dupCnt: 1, sz: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.dupCnt++
	}
	o.maintain()
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.dupCnt > 1 {
			o.dupCnt--
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
	}
	o.maintain()
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }

// < key 的元素个数
func (t *treap) rank(key int) (kth int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			kth += o.lr[0].size() + o.dupCnt
			o = o.lr[1]
		default:
			kth += o.lr[0].size()
			return
		}
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func countRectangles2(rectangles [][]int, points [][]int) (ans []int) {
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][1] > rectangles[j][1] })
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })

	ans = make([]int, len(points))
	n := len(rectangles)
	i := 0
	t := newTreap()
	for _, p := range points {
		for ; i < n && p[1] <= rectangles[i][1]; i++ {
			t.put(rectangles[i][0])
		}
		ans[p[2]] = i - t.rank(p[0])
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][1] > rectangles[j][1] })
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })

	ans := make([]int, len(points))
	i, n, xs := 0, len(rectangles), []int{}
	for _, p := range points {
		start := i
		for ; i < n && p[1] <= rectangles[i][1]; i++ {
			xs = append(xs, rectangles[i][0])
		}
		if start < i { // 只有在 xs 插入了新元素时才排序
			sort.Ints(xs)
		}
		ans[p[2]] = i - sort.SearchInts(xs, p[0])
	}
	return ans
}
