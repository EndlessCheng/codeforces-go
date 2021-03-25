package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type node struct {
	lr       [2]*node
	priority uint
	key      int
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

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key}
	}
	d := o.cmp(key)
	o.lr[d] = t._put(o.lr[d], key)
	if o.lr[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
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
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

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

type trieNode struct {
	son          [26]*trieNode
	allID, curID *treap
}

type trie struct{ root *trieNode }

func (trie) ord(c byte) byte { return c - 'a' }

func (t *trie) put(s []byte, id int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{allID: &treap{rd: 1}, curID: &treap{rd: 1}}
		}
		o = o.son[b]
		o.allID.put(id)
	}
	o.curID.put(id)
}

func (t *trie) delete(s []byte, id int) {
	os := []*trieNode{}
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return
		}
		os = append(os, o)
	}
	o.curID.delete(id)
	for _, o := range os {
		o.allID.delete(id)
	}
}

func (t *trie) hasPrefixOfString(s []byte, l, r int) bool {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return false
		}
		if to := o.curID.lowerBound(l); to != nil && to.key <= r {
			return true
		}
	}
	return false
}

func (t *trie) hasStringOfPrefix(p []byte, l, r int) bool {
	o := t.root
	for _, b := range p {
		o = o.son[t.ord(b)]
		if o == nil {
			return false
		}
	}
	to := o.allID.lowerBound(l)
	return to != nil && to.key <= r
}

func CF101628K(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := &trie{&trieNode{}}
	var n, q, op, i, l, r int
	var s []byte
	Fscan(in, &n)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		t.put(a[i], i+1)
	}
	for Fscan(in, &q); q > 0; q-- {
		switch Fscan(in, &op); op {
		case 1:
			Fscan(in, &i, &s)
			t.delete(a[i-1], i)
			t.put(s, i)
			a[i-1] = s
		case 2:
			Fscan(in, &l, &r, &s)
			if t.hasPrefixOfString(s, l, r) {
				Fprintln(out, "Y")
			} else {
				Fprintln(out, "N")
			}
		default:
			Fscan(in, &l, &r, &s)
			if t.hasStringOfPrefix(s, l, r) {
				Fprintln(out, "Y")
			} else {
				Fprintln(out, "N")
			}
		}
	}
}

//func main() { CF101628K(os.Stdin, os.Stdout) }
