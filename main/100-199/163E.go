package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
)

// https://space.bilibili.com/206214
type acamNode63 struct {
	son  [26]uint32
	fail uint32
}

var acamNodes63 [1e6 + 1]acamNode63

type acam63 struct {
	root    uint32
	nodeCnt uint32
	g       [][]uint32
}

func (t *acam63) put(s string) uint32 {
	o := t.root
	for _, b := range s {
		b -= 'a'
		if acamNodes63[o].son[b] == 0 {
			acamNodes63[o].son[b] = t.nodeCnt
			t.nodeCnt++
		}
		o = acamNodes63[o].son[b]
	}
	return o
}

func (t *acam63) buildFail() {
	t.g = make([][]uint32, t.nodeCnt)
	q := make([]uint32, 0, t.nodeCnt)
	for _, son := range acamNodes63[t.root].son[:] {
		if son != 0 {
			t.g[acamNodes63[son].fail] = append(t.g[acamNodes63[son].fail], son)
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		f := acamNodes63[o].fail
		for i, son := range acamNodes63[o].son[:] {
			if son == 0 {
				acamNodes63[o].son[i] = acamNodes63[f].son[i]
				continue
			}
			acamNodes63[son].fail = acamNodes63[f].son[i]
			t.g[acamNodes63[son].fail] = append(t.g[acamNodes63[son].fail], son)
			q = append(q, son)
		}
	}
}

type fenwick63 []int32

func (f fenwick63) update(i, j uint32, val int32) {
	for ; i < uint32(len(f)); i += i & -i {
		f[i] += val
	}
	for ; j < uint32(len(f)); j += j & -j {
		f[j] -= val
	}
}

func (f fenwick63) pre(i uint32) (res int) {
	for ; i > 0; i &= i - 1 {
		res += int(f[i])
	}
	return
}

func CF163E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, n int
	var s string
	Fscan(in, &q, &n)
	t := &acam63{nodeCnt: 1}
	nodeIDs := make([]uint32, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s)
		nodeIDs[i] = t.put(s)
	}
	t.buildFail()

	g := t.g
	nodes := make([]struct{ l, r uint32 }, len(g))
	type pair struct{ v, i uint32 }

	// 手写递归栈，递归转循环，不然消耗 > 100MB 内存 MLE
	st := []pair{{t.root, 0}}
	nodes[t.root].l = 1
	dfn := uint32(1)
	for len(st) > 0 {
		p := st[len(st)-1]
		v, i := p.v, p.i
		if i < uint32(len(g[v])) {
			dfn++
			w := g[v][i]
			nodes[w].l = dfn
			st[len(st)-1].i++
			st = append(st, pair{w, 0})
		} else {
			nodes[v].r = dfn + 1
			st = st[:len(st)-1]
		}
	}

	bit := make(fenwick63, dfn+2)
	for i := 1; i <= n; i++ {
		p := nodes[nodeIDs[i]]
		bit.update(p.l, p.r, 1)
	}
	del := make([]bool, n+1)
	for ; q > 0; q-- {
		Fscan(in, &s)
		if s[0] == '?' {
			o := t.root
			ans := 0
			for _, b := range s[1:] {
				o = acamNodes63[o].son[b-'a']
				ans += bit.pre(nodes[o].l)
			}
			Fprintln(out, ans)
		} else {
			i, _ := strconv.Atoi(s[1:])
			if del[i] == (s[0] == '-') {
				continue
			}
			del[i] = !del[i]
			p := nodes[nodeIDs[i]]
			bit.update(p.l, p.r, int32(s[0]&3)-2)
		}
	}
}

//func main() { CF163E(os.Stdin, os.Stdout) }
