package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node struct {
	childIdx [2]int
}

type trie struct {
	nodes []*node
}

func (t *trie) insert(s []byte) {
	o := t.nodes[0]
	for _, c := range s {
		if o.childIdx[c] == 0 {
			o.childIdx[c] = len(t.nodes)
			t.nodes = append(t.nodes, &node{})
		}
		o = t.nodes[o.childIdx[c]]
	}
}

func (t *trie) maxXor(bits []byte) (v int) {
	o := t.nodes[0]
	for i, b := range bits {
		if o.childIdx[b^1] > 0 {
			v |= 1 << uint(30-i)
			b ^= 1
		}
		o = t.nodes[o.childIdx[b]]
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	bits32 := func(n int) []byte {
		bits := make([]byte, 31)
		for i := range bits {
			if n>>uint(30-i)&1 == 1 {
				bits[i] = 1
			}
		}
		return bits
	}

	t := &trie{[]*node{{}}}
	var n, v int
	Fscan(in, &n, &v)
	t.insert(bits32(v))
	ans := 0
	for n--; n > 0; n-- {
		Fscan(in, &v)
		bits := bits32(v)
		if v := t.maxXor(bits); v > ans {
			ans = v
		}
		t.insert(bits)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
