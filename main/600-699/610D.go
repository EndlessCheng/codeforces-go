package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type node10 struct {
	ch       [2]*node10
	priority uint
	key, sz  int
}

func (o *node10) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node10) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node10) maintain() { o.sz = 1 + o.ch[0].size() + o.ch[1].size() }

func (o *node10) rotate(d int) *node10 {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap10 struct {
	rd   uint
	root *node10
}

func (t *treap10) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap10) _put(o *node10, key int) *node10 {
	if o == nil {
		return &node10{priority: t.fastRand(), key: key, sz: 1}
	}
	d := o.cmp(key)
	o.ch[d] = t._put(o.ch[d], key)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	o.maintain()
	return o
}

func (t *treap10) put(key int) { t.root = t._put(t.root, key) }

func (t *treap10) _delete(o *node10, key int) *node10 {
	if d := o.cmp(key); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], key)
	} else {
		if o.ch[1] == nil {
			return o.ch[0]
		}
		if o.ch[0] == nil {
			return o.ch[1]
		}
		d = 0
		if o.ch[0].priority > o.ch[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.ch[d] = t._delete(o.ch[d], key)
	}
	o.maintain()
	return o
}

func (t *treap10) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap10) rank(key int) (kth int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.ch[0]
		case c > 0:
			kth += 1 + o.ch[0].size()
			o = o.ch[1]
		default:
			kth += o.ch[0].size()
			return
		}
	}
	return
}

func CF610D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ p, l, r int }
	var a, b []pair
	var n, x1, y1, x2, y2 int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			a = append(a, pair{y1, x1, x2})
		} else {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			b = append(b, pair{x1, y1, y2})
		}
	}

	ans := int64(0)
	unique := func(a []pair) (b []pair) {
		sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.p < b.p || a.p == b.p && a.l < b.l })
		for _, p := range a {
			if b == nil || p.p > b[len(b)-1].p || p.l > b[len(b)-1].r {
				b = append(b, p)
			} else if p.r > b[len(b)-1].r {
				b[len(b)-1].r = p.r // 合并线段
			}
		}
		for _, p := range b {
			ans += int64(p.r - p.l + 1)
		}
		return b
	}
	a = unique(a)
	b = unique(b)
	if len(a) > len(b) {
		a, b = b, a
	}

	type event struct{ e, p int }
	es := make([]event, 0, 2*len(a))
	for _, p := range a {
		es = append(es, event{p.l<<1 | 1, p.p}, event{(p.r + 1) << 1, p.p})
	}
	sort.Slice(es, func(i, j int) bool { return es[i].e < es[j].e })
	t := &treap10{rd: 1} //
	i := 0
	for _, p := range b {
		for ; i < len(es) && es[i].e>>1 <= p.p; i++ {
			if es[i].e&1 > 0 {
				t.put(es[i].p)
			} else {
				t.delete(es[i].p)
			}
		}
		ans -= int64(t.rank(p.r+1) - t.rank(p.l)) // 减去交点个数
	}
	Fprint(out, ans)
}

//func main() { CF610D(os.Stdin, os.Stdout) }
