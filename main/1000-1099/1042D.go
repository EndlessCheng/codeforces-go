package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type node struct {
	lr       [2]*node
	priority uint
	key      int
	keyCnt   int
	subCnt   int
}

func (o *node) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node) maintain() {
	o.subCnt = o.keyCnt + o.lr[0].size() + o.lr[1].size()
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap42 struct {
	rd   uint
	root *node
}

func (t *treap42) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap42) _put(o *node, key int) *node {
	if o == nil {
		o = &node{priority: t.fastRand(), key: key, keyCnt: 1}
	} else if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.keyCnt++
	}
	o.maintain()
	return o
}

func (t *treap42) put(key int) { t.root = t._put(t.root, key) }

func (o *node) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a > b {
		return 0
	}
	return 1
}

func (t *treap42) gr(key int) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			cnt += o.lr[0].size() + o.keyCnt
			o = o.lr[1]
		default:
			cnt += o.lr[0].size()
			return
		}
	}
	return
}

func cf1042D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, upper, v, s, ans int
	Fscan(in, &n, &upper)
	t := &treap42{rd: 1}
	t.put(0)
	for ; n > 0; n-- {
		Fscan(in, &v)
		s += v
		ans += t.gr(s - upper)
		t.put(s)
	}
	Fprint(out, ans)
}

//func main() { cf1042D(os.Stdin, os.Stdout) }
