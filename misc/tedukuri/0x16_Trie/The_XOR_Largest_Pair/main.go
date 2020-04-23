package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node struct{ son [2]*node }
type trie struct{ root *node }

func newTrie() *trie {
	// init with a root (empty string)
	return &trie{&node{}}
}

func (t *trie) put(val int) {
	s := [31]byte{}
	for i := range s {
		s[i] = byte(val >> (30 - i) & 1)
	}
	o := t.root
	for _, c := range s {
		if o.son[c] == nil {
			o.son[c] = &node{}
		}
		o = o.son[c]
	}
}

func (t *trie) maxXor(val int) (ans int) {
	bits := [31]byte{}
	for i := range bits {
		bits[i] = byte(val >> (30 - i) & 1)
	}
	o := t.root
	for i, b := range bits {
		if o.son[b^1] != nil {
			ans |= 1 << (30 - i)
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := newTrie()
	var n, v, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		t.put(v)
		if v = t.maxXor(v); v > ans {
			ans = v
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
