package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type acamNode struct {
	son    [26]*acamNode
	fail   *acamNode
	fa     *acamNode
	nodeID int
}

type gInfo struct{ l, r int }

type acam struct {
	root    *acamNode
	nodeCnt int

	g     [][]int
	gInfo []gInfo
	dfn   int
}

func (t *acam) addEdge(v, w int) { t.g[v] = append(t.g[v], w) }

func (t *acam) put(s string) *acamNode {
	o := t.root
	for _, b := range s {
		b -= 'a'
		if o.son[b] == nil {
			o.son[b] = &acamNode{fa: o, nodeID: t.nodeCnt}
			t.nodeCnt++
		}
		o = o.son[b]
	}
	return o
}

func (t *acam) buildFail() {
	t.g = make([][]int, t.nodeCnt)
	t.root.fail = t.root
	q := make([]*acamNode, 0, t.nodeCnt)
	for i, son := range t.root.son[:] {
		if son == nil {
			t.root.son[i] = t.root
		} else {
			son.fail = t.root
			t.addEdge(son.fail.nodeID, son.nodeID)
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		f := o.fail
		for i, son := range o.son[:] {
			if son == nil {
				o.son[i] = f.son[i]
				continue
			}
			son.fail = f.son[i]
			t.addEdge(son.fail.nodeID, son.nodeID)
			q = append(q, son)
		}
	}
}

func (t *acam) buildDFN(v int) {
	t.dfn++
	t.gInfo[v].l = t.dfn
	for _, w := range t.g[v] {
		t.buildDFN(w)
	}
	t.gInfo[v].r = t.dfn
}

type fenwick47 []int

func (f fenwick47) update(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick47) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick47) query(l, r int) (res int) {
	return f.pre(r) - f.pre(l-1)
}

func CF547E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := &acam{
		root:    &acamNode{},
		nodeCnt: 1,
	}
	var n, q, l, r, k int
	var s string
	Fscan(in, &n, &q)
	a := make([]*acamNode, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s)
		a[i] = t.put(s)
	}
	t.buildFail()

	t.gInfo = make([]gInfo, len(t.g))
	t.buildDFN(t.root.nodeID)

	type query struct{ qid, k, sgn int }
	qs := make([][]query, n+1)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r, &k)
		qs[l-1] = append(qs[l-1], query{i, k, -1})
		qs[r] = append(qs[r], query{i, k, 1})
	}

	ans := make([]int, q)
	bit := make(fenwick47, t.nodeCnt+1)
	for i := 1; i <= n; i++ {
		for o := a[i]; o != t.root; o = o.fa {
			bit.update(t.gInfo[o.nodeID].l)
		}
		for _, q := range qs[i] {
			p := t.gInfo[a[q.k].nodeID]
			ans[q.qid] += q.sgn * bit.query(p.l, p.r)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF547E(os.Stdin, os.Stdout) }
