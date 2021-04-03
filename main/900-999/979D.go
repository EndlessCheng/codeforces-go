package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
)

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type node79 struct {
	son [2]*node79
	min int
}

type trie79 struct{ root *node79 }

func (t *trie79) put(v int) *node79 {
	o := t.root
	if v < o.min {
		o.min = v
	}
	for i := 16; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &node79{min: 2e9}
		}
		o = o.son[b]
		if v < o.min {
			o.min = v
		}
	}
	return o
}

func (t *trie79) maxXorWithLimitVal(v, limit int) (w int) {
	o := t.root
	if o.min > limit {
		return -1
	}
	for i := 16; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil && o.son[b^1].min <= limit {
			b ^= 1
		}
		w |= b << i
		o = o.son[b]
	}
	return
}

func CF979D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e5
	ds := [mx + 1][]int{}
	ts := [mx + 1]*trie79{}
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			ds[j] = append(ds[j], i)
		}
		ts[i] = &trie79{&node79{min: 2e9}}
	}
	has := [mx + 1]bool{}

	var q, t, x, k, s int
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &t, &x); t == 1 {
			if !has[x] {
				has[x] = true
				for _, d := range ds[x] {
					ts[d].put(x)
				}
			}
		} else if Fscan(in, &k, &s); x%k > 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ts[k].maxXorWithLimitVal(x, s-x))
		}
	}
}

//func main() { CF979D(os.Stdin, os.Stdout) }
