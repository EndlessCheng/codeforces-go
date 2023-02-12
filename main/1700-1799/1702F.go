package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"time"
)

// https://space.bilibili.com/206214
type node02 struct {
	lr       [2]*node02
	priority uint
	key      int
	keyCnt   int
	subCnt   int
}

func (o *node02) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node02) maintain() {
	o.subCnt = o.keyCnt + o.lr[0].size() + o.lr[1].size()
}

func (o *node02) rotate(d int) *node02 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap02 struct {
	rd   uint
	root *node02
}

func (t *treap02) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap02) size() int { return t.root.size() }

func (t *treap02) _put(o *node02, key, c int) *node02 {
	if o == nil {
		o = &node02{priority: t.fastRand(), key: key, keyCnt: c}
	} else if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, c)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.keyCnt += c
	}
	o.maintain()
	return o
}

func (t *treap02) put(key, c int) { t.root = t._put(t.root, key, c) }

func (t *treap02) _delete(o *node02, key, c int) *node02 {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key, c)
	} else {
		if o.keyCnt > c {
			o.keyCnt -= c
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
			o.lr[d] = t._delete(o.lr[d], key, c)
		}
	}
	o.maintain()
	return o
}

func (t *treap02) delete(key, c int) { t.root = t._delete(t.root, key, c) }

func newTreap() *treap02 { return &treap02{rd: uint(time.Now().UnixNano())/2 + 1} }

func (o *node02) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap02) max() (max *node02) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

func CF1702F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := newTreap()
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			a.put(v>>bits.TrailingZeros(uint(v)), 1)
		}
		b := newTreap()
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			b.put(v>>bits.TrailingZeros(uint(v)), 1)
		}

		for a.size() > 0 {
			p := b.max()
			q := a.max()
			if p.key < q.key {
				Fprintln(out, "NO")
				continue o
			}
			if p.key == q.key {
				if p.keyCnt < q.keyCnt {
					Fprintln(out, "NO")
					continue o
				}
				a.delete(q.key, q.keyCnt)
				b.delete(p.key, p.keyCnt)
				if p.keyCnt > q.keyCnt {
					b.put(p.key>>1, p.keyCnt-q.keyCnt)
				}
			} else {
				b.delete(p.key, p.keyCnt)
				b.put(p.key>>1, p.keyCnt)
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1702F(os.Stdin, os.Stdout) }
