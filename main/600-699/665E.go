package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
)

// https://space.bilibili.com/206214
func init() { debug.SetGCPercent(-1) }

type trieNode65 struct {
	son [2]*trieNode65
	cnt int
}

type trie65 struct{ root *trieNode65 }

func (t *trie65) put(v int) *trieNode65 {
	o := t.root
	for i := 29; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode65{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie65) countLimitXOR(v, limit int) (cnt int) {
	o := t.root
	for i := 29; i >= 0; i-- {
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

func CF665E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, xor int
	Fscan(in, &n, &k)
	ans := int64(n) * int64(n+1) / 2
	t := &trie65{&trieNode65{}}
	t.put(0)
	for ; n > 0; n-- {
		Fscan(in, &v)
		xor ^= v
		ans -= int64(t.countLimitXOR(xor, k))
		t.put(xor)
	}
	Fprint(out, ans)
}

//func main() { CF665E(os.Stdin, os.Stdout) }
