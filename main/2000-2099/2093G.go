package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type node93 struct {
	son [2]*node93
	cnt int
}

type trie93 struct{ root *node93 }

const trieBitLen93 = 30

func (t *trie93) put(v int) *node93 {
	o := t.root
	for i := trieBitLen93 - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &node93{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie93) del(v int) *node93 {
	o := t.root
	for i := trieBitLen93 - 1; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.cnt--
	}
	return o
}

func (t *trie93) maxXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen93 - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil && o.son[b^1].cnt > 0 {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

func cf2093G(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if k == 0 {
			Fprintln(out, 1)
			continue
		}
		t := &trie93{&node93{}}
		ans, left := n+1, 0
		for i, v := range a {
			t.put(v)
			for t.maxXor(v) >= k {
				ans = min(ans, i-left+1)
				t.del(a[left])
				left++
			}
		}
		if ans > n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2093G(bufio.NewReader(os.Stdin), os.Stdout) }
