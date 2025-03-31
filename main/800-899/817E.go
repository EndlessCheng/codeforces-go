package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
)

// https://github.com/EndlessCheng
func init() { debug.SetGCPercent(-1) }

type node struct {
	son [2]*node
	cnt int
}

type trie struct{ root *node }

const trieBitLen = 27

func (t *trie) put(v int) *node {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &node{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie) del(v int) *node {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.cnt--
	}
	return o
}

func (t *trie) countLimitXOR(v, limit int) (cnt int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				cnt += o.son[b].cnt
			}
			b ^= 1
		}
		if o.son[b] == nil {
			return
		}
		o = o.son[b]
	}
	return
}

func cf817E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, op, v, l int
	t := trie{&node{}}
	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &v)
		if op == 1 {
			t.put(v)
		} else if op == 2 {
			t.del(v)
		} else {
			Fscan(in, &l)
			Fprintln(out, t.countLimitXOR(v, l))
		}
	}
}

//func main() { cf817E(bufio.NewReader(os.Stdin), os.Stdout) }
