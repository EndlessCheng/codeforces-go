package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
type node88 struct{ son [2]*node88 }
type trie88 struct{ root *node88 }

const trieBitLen88 = 30

func (t *trie88) put(v int32) *node88 {
	o := t.root
	for i := trieBitLen88 - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &node88{}
		}
		o = o.son[b]
	}
	return o
}

func (t *trie88) minXor(v int32) (ans int) {
	o := t.root
	for i := trieBitLen88 - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

func CF888G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int32, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	var f func([]int32, int)
	f = func(a []int32, p int) {
		if a == nil || p < 0 {
			return
		}
		b := [2][]int32{}
		for _, v := range a {
			k := v >> p & 1
			b[k] = append(b[k], v)
		}
		if b[0] != nil && b[1] != nil {
			t := &trie88{&node88{}}
			for _, v := range b[0] {
				t.put(v)
			}
			res := math.MaxInt
			for _, v := range b[1] {
				res = min(res, t.minXor(v))
			}
			ans += res
		}
		f(b[0], p-1)
		f(b[1], p-1)
	}
	f(a, 29)
	Fprint(out, ans)
}

//func main() { CF888G(os.Stdin, os.Stdout) }
