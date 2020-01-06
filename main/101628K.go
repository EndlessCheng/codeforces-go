package main

import (
	"bufio"
	. "fmt"
	"io"
)

type sNode struct {
	lr       [2]*sNode
	priority uint
	key      int
}

func (o *sNode) rotate(d int) *sNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type sTreap struct {
	rd   uint
	root *sNode
}

func (t *sTreap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *sTreap) compare(a, b int) int {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *sTreap) _put(o *sNode, key int) *sNode {
	if o == nil {
		return &sNode{priority: t.fastRand(), key: key}
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	}
	return o
}

func (t *sTreap) put(key int) { t.root = t._put(t.root, key) }

func (t *sTreap) _delete(o *sNode, key int) *sNode {
	if o == nil {
		return nil
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		cmp2 := 0
		if o.lr[0].priority > o.lr[1].priority {
			cmp2 = 1
		}
		o = o.rotate(cmp2)
		o.lr[cmp2] = t._delete(o.lr[cmp2], key)
	}
	return o
}

func (t *sTreap) delete(key int) { t.root = t._delete(t.root, key) }

func (t *sTreap) ceiling(key int) (ceiling *sNode) {
	for o := t.root; o != nil; {
		switch cmp := t.compare(key, o.key); {
		case cmp == 0:
			ceiling = o
			o = o.lr[0]
		case cmp > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *sTreap) hasValueInRange(l, r int) bool {
	o := t.ceiling(l)
	return o != nil && o.key <= r
}

type trieNode struct {
	childIdx       [26]int
	curIndexes     *sTreap
	subTreeIndexes *sTreap
}

type trie struct {
	nodes []*trieNode
}

func (t *trie) put(s string, idx int) {
	o := t.nodes[0]
	for i := range s {
		c := s[i] - 'a'
		if o.childIdx[c] == 0 {
			o.childIdx[c] = len(t.nodes)
			t.nodes = append(t.nodes, &trieNode{
				curIndexes:     &sTreap{rd: 1},
				subTreeIndexes: &sTreap{rd: 1},
			})
		}
		o = t.nodes[o.childIdx[c]]
		o.subTreeIndexes.put(idx)
	}
	o.curIndexes.put(idx)
}

func (t *trie) del(s string, idx int) {
	o := t.nodes[0]
	for i := range s {
		o = t.nodes[o.childIdx[s[i]-'a']]
		o.subTreeIndexes.delete(idx)
	}
	o.curIndexes.delete(idx)
}

func (t *trie) hasPrefixOfText(s string, l, r int) bool {
	o := t.nodes[0]
	for i := range s {
		idx := o.childIdx[s[i]-'a']
		if idx == 0 {
			return false
		}
		o = t.nodes[idx]
		if o.curIndexes.hasValueInRange(l, r) {
			return true
		}
	}
	return false
}

func (t *trie) hasTextOfPrefix(p string, l, r int) bool {
	o := t.nodes[0]
	for i := range p {
		idx := o.childIdx[p[i]-'a']
		if idx == 0 {
			return false
		}
		o = t.nodes[idx]
	}
	return o.subTreeIndexes.hasValueInRange(l, r)
}

// github.com/EndlessCheng/codeforces-go
func Sol101628K(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, idx, l, r int
	var s string
	Fscan(in, &n)
	t := &trie{nodes: []*trieNode{{}}}
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
		t.put(a[i], i+1)
	}
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		Fscan(in, &op)
		switch op {
		case 1:
			Fscan(in, &idx, &s)
			t.del(a[idx-1], idx)
			t.put(s, idx)
			a[idx-1] = s
		case 2:
			Fscan(in, &l, &r, &s)
			if t.hasPrefixOfText(s, l, r) {
				Fprintln(out, "Y")
			} else {
				Fprintln(out, "N")
			}
		default:
			Fscan(in, &l, &r, &s)
			if t.hasTextOfPrefix(s, l, r) {
				Fprintln(out, "Y")
			} else {
				Fprintln(out, "N")
			}
		}
	}
}

//func main() {
//	Sol101628K(os.Stdin, os.Stdout)
//}
