package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node38 struct {
	lr       [2]*node38
	priority uint
	key      int
	subCnt   int
}

func (o *node38) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node38) maintain() { o.subCnt = 1 + o.lr[0].size() + o.lr[1].size() }

func (o *node38) rotate(d int) *node38 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap38 struct {
	rd   uint
	root *node38
}

func (t *treap38) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap38) size() int { return t.root.size() }

func (t *treap38) _put(o *node38, key int) *node38 {
	if o == nil {
		return &node38{priority: t.fastRand(), key: key, subCnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		// 相等
	}
	o.maintain()
	return o
}

func (t *treap38) put(key int) { t.root = t._put(t.root, key) }

func (t *treap38) _delete(o *node38, key int) *node38 {
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

func (t *treap38) delete(key int) { t.root = t._delete(t.root, key) }

func (o *node38) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap38) min() (min int) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o.key
	}
	return
}

func (t *treap38) max() (max int) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o.key
	}
	return
}

func CF1838D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, p int
	var s []byte
	Fscan(in, &n, &q, &s)
	t := &treap38{rd: uint(time.Now().UnixNano())/2 + 1}
	for i, b := range s {
		if b%2 != byte(i%2) {
			t.put(i)
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &p)
		p--
		s[p] ^= 1
		if s[p]%2 != byte(p%2) {
			t.put(p)
		} else {
			t.delete(p)
		}
		if n%2 > 0 || t.size() > 0 && t.min()%2 <= t.max()%2 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1838D(os.Stdin, os.Stdout) }
