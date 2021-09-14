package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type node struct {
	lr          [2]*node
	next        *node
	priority    uint
	pos, spd, i int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.pos:
		return 0
	case b > o.pos:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
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

func (t *treap) _put(o *node, key, val, i int, next *node) *node {
	if o == nil {
		return &node{priority: t.fastRand(), pos: key, spd: val, i: i, next: next}
	}
	d := o.cmp(key)
	o.lr[d] = t._put(o.lr[d], key, val, i, next)
	if o.lr[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *treap) put(key, val, i int, next *node) { t.root = t._put(t.root, key, val, i, next) }

func (t *treap) lowerBound(key int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func newTreap() *treap { return &treap{rd: 1} }

func getCollisionTimes(a [][]int) []float64 {
	n := len(a)
	ans2 := make([][2]int, n)

	for i, p := range a {
		a[i] = append(p, i)
	}
	sort.Slice(a, func(i, j int) bool { return a[i][1] < a[j][1] })

	t := newTreap()
	for i := 0; i < n; {
		st := i
		groupSpd := a[st][1]
		nexts := []*node{}
		for ; i < n && a[i][1] == groupSpd; i++ {
			curPos, curSpd, curI := a[i][0], a[i][1], a[i][2]
			o := t.lowerBound(curPos)
			if o == nil {
				ans2[curI][0] = -1
				nexts = append(nexts, nil)
				continue
			}
			for {
				dx := o.pos - curPos
				ds := curSpd - o.spd
				if o.next == nil || dx*ans2[o.i][1] < ds*ans2[o.i][0] {
					ans2[curI] = [2]int{dx, ds}
					nexts = append(nexts, o)
					break
				}
				o = o.next
			}
		}
		for j, p := range a[st:i] {
			t.put(p[0], p[1], p[2], nexts[j])
		}
	}

	ans := make([]float64, n)
	for i, p := range ans2 {
		if p[0] == -1 {
			ans[i] = -1
		} else {
			ans[i] = float64(p[0]) / float64(p[1])
		}
	}
	return ans
}
