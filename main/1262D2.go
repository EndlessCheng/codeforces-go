package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

type pstNode1262D2 struct {
	l, r   int
	lo, ro *pstNode1262D2
	sum    int
}
type pSegmentTree1262D2 struct {
	nodes        []pstNode1262D2
	versionRoots []*pstNode1262D2
}

func newPST1262D2(n int) *pSegmentTree1262D2 {
	maxNodeSize := n * (3 + bits.Len(uint(n)))
	return &pSegmentTree1262D2{
		make([]pstNode1262D2, 0, maxNodeSize),
		make([]*pstNode1262D2, 1, n+1),
	}
}

func (t *pSegmentTree1262D2) _build(l, r int) *pstNode1262D2 {
	t.nodes = append(t.nodes, pstNode1262D2{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		return o
	}
	mid := (l + r) >> 1
	o.lo = t._build(l, mid)
	o.ro = t._build(mid+1, r)
	return o
}

func (t *pSegmentTree1262D2) _update(o *pstNode1262D2, idx int) *pstNode1262D2 {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.sum++
		return o
	}
	if idx <= o.lo.r {
		o.lo = t._update(o.lo, idx)
	} else {
		o.ro = t._update(o.ro, idx)
	}
	o.sum = o.lo.sum + o.ro.sum
	return o
}

func (t *pSegmentTree1262D2) _queryKth(o1, o2 *pstNode1262D2, k int) (idx int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := o2.lo.sum - o1.lo.sum; d >= k {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-d)
	}
}

func (t *pSegmentTree1262D2) init(n int) {
	t.versionRoots[0] = t._build(1, n)
}

func (t *pSegmentTree1262D2) update(baseVersion int, idx int) {
	t.versionRoots = append(t.versionRoots, t._update(t.versionRoots[baseVersion], idx))
}

func (t *pSegmentTree1262D2) queryKth(l, r int, k int) (idx int) {
	return t._queryKth(t.versionRoots[l-1], t.versionRoots[r], k)
}

// github.com/EndlessCheng/codeforces-go
func Sol1262D2(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	a := make([]int, n)
	type pair struct{ val, i int }
	ps := make([]pair, n)
	for i := range ps {
		a[i] = read()
		ps[i] = pair{a[i], i}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].val > ps[j].val || ps[i].val == ps[j].val && ps[i].i < ps[j].i
	})
	st := newPST1262D2(n)
	st.init(n)
	for i, p := range ps {
		st.update(i, p.i+1)
	}

	for m := read(); m > 0; m-- {
		idx := st.queryKth(1, read(), read()) - 1
		Fprintln(out, a[idx])
	}
}

//func main() {
//	Sol1262D2(os.Stdin, os.Stdout)
//}
