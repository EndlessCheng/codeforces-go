package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

type pstNode struct {
	l, r   int
	lo, ro *pstNode
	val    int
}
type pst struct {
	nodes        []pstNode
	versionRoots []*pstNode
}

func (t *pst) _buildArr(arr []int, l, r int) *pstNode {
	t.nodes = append(t.nodes, pstNode{l: l, r: r})
	o := &t.nodes[len(t.nodes)-1]
	if l == r {
		o.val = arr[l]
		return o
	}
	mid := (l + r) >> 1
	o.lo = t._buildArr(arr, l, mid)
	o.ro = t._buildArr(arr, mid+1, r)
	return o
}

func (t *pst) _update(o *pstNode, idx int, val int) *pstNode {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.val = val
		return o
	}
	if mid := o.lo.r; idx <= mid {
		o.lo = t._update(o.lo, idx, val)
	} else {
		o.ro = t._update(o.ro, idx, val)
	}
	return o
}

func (t *pst) _query(o *pstNode, idx int) (res int) {
	if o.l == o.r {
		return o.val
	}
	mid := o.lo.r
	if idx <= mid {
		return t._query(o.lo, idx)
	}
	return t._query(o.ro, idx)
}

func (t *pst) initArr(arr []int) {
	t.versionRoots[0] = t._buildArr(arr, 1, len(arr)-1)
}

func (t *pst) copy(dstVersion, srcVersion int) {
	t.versionRoots[dstVersion] = t.versionRoots[srcVersion]
}

func (t *pst) update(dstVersion, srcVersion int, idx int, val int) {
	t.versionRoots[dstVersion] = t._update(t.versionRoots[srcVersion], idx, val)
}

func (t *pst) query(version int, idx int) (sum int) {
	return t._query(t.versionRoots[version], idx)
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

	n, m := read(), read()
	t := &pst{
		make([]pstNode, 0, 2*n+(bits.Len(uint(n))+1)*m),
		make([]*pstNode, m+1),
	}
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = read()
	}
	t.initArr(a)
	for i := 1; i <= m; i++ {
		ver, op, pos := read(), read(), read()
		if op == 1 {
			t.update(i, ver, pos, read())
		} else {
			Fprintln(out, t.query(ver, pos))
			t.copy(i, ver)
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
