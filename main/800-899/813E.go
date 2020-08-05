package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

type pstNode813E struct {
	l, r   int
	lo, ro *pstNode813E
	sum    int
}
type pst813E struct {
	nodes        []pstNode813E
	versionRoots []*pstNode813E
}

func (t *pst813E) _build(l, r int) *pstNode813E {
	t.nodes = append(t.nodes, pstNode813E{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		return o
	}
	mid := (l + r) >> 1
	o.lo = t._build(l, mid)
	o.ro = t._build(mid+1, r)
	return o
}

func (t *pst813E) _update(o *pstNode813E, idx int, val int) *pstNode813E {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.sum += val
		return o
	}
	if mid := o.lo.r; idx <= mid {
		o.lo = t._update(o.lo, idx, val)
	} else {
		o.ro = t._update(o.ro, idx, val)
	}
	o.sum = o.lo.sum + o.ro.sum
	return o
}

func (t *pst813E) _query(o *pstNode813E, l, r int) (res int) {
	if l <= o.l && o.r <= r {
		return o.sum
	}
	mid := o.lo.r
	if l <= mid {
		res += t._query(o.lo, l, r)
	}
	if mid < r {
		res += t._query(o.ro, l, r)
	}
	return
}

func (t *pst813E) init(n int) {
	t.versionRoots[0] = t._build(1, n)
}

func (t *pst813E) update(dstVersion, srcVersion int, idx int, val int) {
	t.versionRoots[dstVersion] = t._update(t.versionRoots[srcVersion], idx, val)
}

func (t *pst813E) query(version int, l, r int) (sum int) {
	return t._query(t.versionRoots[version], l, r)
}

// github.com/EndlessCheng/codeforces-go
func Sol813E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k, v, q, l, r int
	Fscan(in, &n, &k)
	t := &pst813E{
		make([]pstNode813E, 0, (bits.Len(uint(n))+2)*2*n),
		make([]*pstNode813E, n+1),
	}
	t.init(n)
	idx := [100001][]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		t.update(i, i-1, i, 1) // 差分
		idx[v] = append(idx[v], i)
		if sz := len(idx[v]); sz > k {
			t.update(i, i, idx[v][sz-k-1], -1) // 差分
		}
	}

	last := 0
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		l = (l+last)%n + 1
		r = (r+last)%n + 1
		if l > r {
			l, r = r, l
		}
		last = t.query(r, l, r)
		Fprintln(out, last)
	}
}

//func main() {
//	Sol813E(os.Stdin, os.Stdout)
//}
