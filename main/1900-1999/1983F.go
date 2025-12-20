package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
const w83 = 31

var nodes83 [1e5 * (w83 + 3 - 17)]node83
var cur83 uint

type node83 struct {
	son [2]*node83
	mx  int
}

func newNode83() *node83 {
	cur83++
	nodes83[cur83] = node83{}
	return &nodes83[cur83]
}

type trie83 struct{ root *node83 }

func (t *trie83) put(v, idx int) *node83 {
	o := t.root
	for i := w83 - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = newNode83()
		}
		o = o.son[b]
		o.mx = idx
	}
	return o
}

func (t *trie83) maxIdx(v, high int) int {
	high++
	res := -1
	o := t.root
	for i := w83 - 1; i >= 0; i-- {
		b := v >> i & 1
		if high>>i&1 > 0 {
			if o.son[b] != nil {
				res = max(res, o.son[b].mx)
			}
			b ^= 1
		}
		if o.son[b] == nil {
			break
		}
		o = o.son[b]
	}
	return res
}

func cf1983F(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := sort.Search(1<<w83, func(high int) bool {
			cur83 = 0
			t := &trie83{newNode83()}
			mxI := -1
			cnt := 0
			for i, v := range a {
				mxI = max(mxI, t.maxIdx(v, high))
				cnt += mxI + 1
				t.put(v, i)
			}
			return cnt >= k
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1983F(bufio.NewReader(os.Stdin), os.Stdout) }
