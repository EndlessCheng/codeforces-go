package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type node49 struct {
	lr       [2]*node49
	priority uint
	key      int
	ids      []int
}

func (o *node49) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node49) rotate(d int) *node49 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap49 struct {
	rd   uint
	root *node49
}

func (t *treap49) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap49) _put(o *node49, key, id int) *node49 {
	if o == nil {
		return &node49{priority: t.fastRand(), key: key, ids: []int{id}}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, id)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.ids = append(o.ids, id)
	}
	return o
}

func (t *treap49) put(key, id int) { t.root = t._put(t.root, key, id) }

func (t *treap49) _delete(o *node49, key int) *node49 {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if len(o.ids) > 1 {
			o.ids = o.ids[:len(o.ids)-1]
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
	return o
}

func (t *treap49) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap49) min() (min *node49) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func (t *treap49) max() (max *node49) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

func CF1249D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, l, r, sz int
	type pair struct{ v, i int }
	rs := make([][]pair, 2e5+1)
	Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		Fscan(in, &l, &r)
		rs[l] = append(rs[l], pair{r, i})
	}

	ans := []interface{}{}
	t := &treap49{rd: 1}
	for l, rs := range rs {
		for {
			o := t.min()
			if o == nil || o.key >= l {
				break
			}
			t.delete(o.key)
			sz--
		}
		for _, p := range rs {
			t.put(p.v, p.i)
			sz++
		}
		for ; sz > k; sz-- {
			o := t.max()
			ans = append(ans, o.ids[len(o.ids)-1])
			o.ids = o.ids[:len(o.ids)-1]
			if len(o.ids) == 0 {
				t.delete(o.key)
			}
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { CF1249D2(os.Stdin, os.Stdout) }
