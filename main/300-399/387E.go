package main

import (
	. "fmt"
	"io"
	"runtime/debug"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type node87 struct {
	lr       [2]*node87
	priority uint
	key      int
}

func (o *node87) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node87) rotate(d int) *node87 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap87 struct {
	rd   uint
	root *node87
}

func (t *treap87) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap87) _put(o *node87, key int) *node87 {
	if o == nil {
		return &node87{priority: t.fastRand(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap87) put(key int) { t.root = t._put(t.root, key) }

func (t *treap87) prev(key int) (prev *node87) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

func (t *treap87) next(key int) (next *node87) {
	for o := t.root; o != nil; {
		if o.cmp(key) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func CF387E(_r io.Reader, out io.Writer) {
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m := r(), r()
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pos[r()] = i
	}
	save := make([]bool, n+1)
	for ; m > 0; m-- {
		save[r()] = true
	}
	tree := make([]int, n+1)
	add := func(i int) {
		for ; i <= n; i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int) (s int) {
		for ; i > 0; i &= i - 1 {
			s += tree[i]
		}
		return
	}

	ans := int64(0)
	t := &treap87{rd: uint(time.Now().UnixNano())/2 + 1}
	t.put(0)
	t.put(n + 1)
	for i := 1; i <= n; i++ {
		p := pos[i]
		if save[i] {
			t.put(p)
		} else {
			l, r := t.prev(p).key, t.next(p).key-1
			ans += int64(r - l - sum(r) + sum(l))
			add(p)
		}
	}
	Fprint(out, ans)
}

//func main() { CF387E(os.Stdin, os.Stdout) }
