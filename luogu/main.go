package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"sort"
)

type pstNode struct {
	l, r   int
	lo, ro *pstNode // 由于使用了指针，pstNode 必须存放于一个分配了足够空间的 slice 中，避免扩容时改变了内存位置
	sum    int
}
type pst struct {
	nodes        []pstNode
	versionRoots []*pstNode
}

func (t *pst) _build(l, r int) *pstNode {
	t.nodes = append(t.nodes, pstNode{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		return o
	}
	mid := (l + r) >> 1
	o.lo = t._build(l, mid)
	o.ro = t._build(mid+1, r)
	return o
}

func (t *pst) _update(o *pstNode, idx int) *pstNode {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.sum++
		return o
	}
	if mid := o.lo.r; idx <= mid {
		o.lo = t._update(o.lo, idx)
	} else {
		o.ro = t._update(o.ro, idx)
	}
	o.sum = o.lo.sum + o.ro.sum
	return o
}

func (t *pst) _queryKth(o1, o2 *pstNode, k int) (allKth int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := o2.lo.sum - o1.lo.sum; d >= k {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-d)
	}
}

func (t *pst) init(n int) {
	t.versionRoots[0] = t._build(1, n)
}

func (t *pst) update(dstVersion, srcVersion int, idx int) {
	t.versionRoots[dstVersion] = t._update(t.versionRoots[srcVersion], idx)
}

func (t *pst) queryKth(l, r int, k int) (allKth int) {
	return t._queryKth(t.versionRoots[l-1], t.versionRoots[r], k)
}

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		data := in.Bytes()
		sig := 1
		if data[0] == '-' {
			sig = -1
			data = data[1:]
		}
		for _, b := range data {
			x = x*10 + int(b-'0')
		}
		return sig * x
	}
	discrete := func(arr []int) (kthArr []int) {
		n := len(arr)
		type pair struct{ val, i int }
		ps := make([]pair, n)
		for i, v := range arr {
			ps[i] = pair{v, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].val < ps[j].val })
		kthArr = make([]int, n)
		for i, p := range ps {
			kthArr[p.i] = i + 1
		}
		return
	}

	n, m := read(), read()
	t := &pst{
		make([]pstNode, 0, (bits.Len(uint(n))+3)*n),
		make([]*pstNode, n+1),
	}
	t.init(n)
	a := make([]int, n)
	for i := range a {
		a[i] = read()
	}
	sortedArr := make([]int, n)
	copy(sortedArr, a)
	sort.Ints(sortedArr)

	kthArr := discrete(a)
	for i, kth := range kthArr {
		t.update(i+1, i, kth)
	}
	for ; m > 0; m-- {
		l, r, k := read(), read(), read()
		Fprintln(out, sortedArr[t.queryKth(l, r, k)-1])
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
