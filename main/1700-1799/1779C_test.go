package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/1779/C
// https://codeforces.com/problemset/status/1779/problem/C
func TestCF1779C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 3
-1 -2 -3 -4
4 3
1 2 3 4
1 1
1
5 5
-2 3 -5 1 -20
5 2
-2 3 -5 -5 -20
10 4
345875723 -48 384678321 -375635768 -35867853 -35863586 -358683842 -81725678 38576 -357865873
outputCopy
1
1
0
0
3
4
inputCopy
1
4 4
-1 5 3 -1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1779C)
}

func TestCompare79C(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 9)
		rg.Int(1,n)
		rg.NewLine()
		rg.IntSlice(n, -5, 5)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, sol79c, CF1779C)
}

type SegmentTree[V, T any] struct {
	initializer func(slice []V, index int) T
	operation   func(lhs, rhs T) T
	tree        []T
}

// constructor
func NewSegmentTree[V, T any](
	initializer func(slice []V, index int) T,
	operation func(lhs, rhs T) T,
) SegmentTree[V, T] {
	return SegmentTree[V, T]{initializer, operation, nil}
}

// build
func (segmentTree *SegmentTree[V, T]) Build(slice []V) {
	paddedSize := int(1)
	for ; paddedSize < len(slice); paddedSize <<= 1 {
	}
	tree := make([]T, paddedSize<<1)
	for i := 0; i < len(slice); i++ { // initialize tree leafs
		tree[i+paddedSize] = segmentTree.initializer(slice, i)
	}
	for i := paddedSize - 1; i >= 1; i-- { // compute internal segments values
		child := i << 1
		tree[i] = segmentTree.operation(tree[child], tree[child+1])
	}
	segmentTree.tree = tree
}

// update
func (segmentTree *SegmentTree[V, T]) Update(slice []V, index int) {
	tree := segmentTree.tree
	address := index + len(tree)>>1
	tree[address] = segmentTree.initializer(slice, index)
	for i := address >> 1; i >= 1; i >>= 1 {
		child := i << 1
		tree[i] = segmentTree.operation(tree[child], tree[child+1])
	}
}

// query
func (segmentTree *SegmentTree[V, T]) Query(start, end int) T {
	tree := segmentTree.tree
	pad := len(tree) >> 1

	lefts, rights := make([]int, 0), make([]int, 0)
	for i, j := start+pad, end+pad; i <= j; i, j = (i+1)>>1, (j-1)>>1 {
		if i%2 == 1 {
			lefts = append(lefts, i)
		}
		if j%2 == 0 {
			rights = append(rights, j)
		}
	}
	sort.Ints(rights)

	ops := append(lefts, rights...)
	res := tree[ops[0]]
	for i := 1; i < len(ops); i++ {
		res = segmentTree.operation(res, tree[ops[i]])
	}
	return res
}

func sol79c(in io.Reader, out io.Writer) {
	var tc int
	fmt.Fscan(in, &tc)

	type MinMax struct {
		min    int
		max    int
		argMin int
		argMax int
	}

	st := NewSegmentTree[int, MinMax](
		func(slice []int, index int) MinMax {
			return MinMax{slice[index], slice[index], index, index}
		},
		func(l, r MinMax) MinMax {
			val := [2]MinMax{l, r}
			minI, maxI := 0, 0
			if l.min > r.min {
				minI = 1
			}
			if l.max < r.max {
				maxI = 1
			}
			return MinMax{val[minI].min, val[maxI].max, val[minI].argMin, val[maxI].argMax}
		},
	)

	for ; tc > 0; tc-- {
		var n, m int
		fmt.Fscan(in, &n, &m)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		prefixSums := make([]int, n)
		cumSum := 0
		for i := 0; i < n; i++ {
			prefixSums[i] = a[i] + cumSum
			cumSum += a[i]
		}

		st.Build(a)
		delta, res := 0, 0
		for i := m - 2; i >= 0; i-- {
			q := st.Query(i+1, m-1)
			if prefixSums[i] < prefixSums[m-1]-delta {
				delta += 2 * q.max
				a[q.argMax] = -a[q.argMax]
				st.Update(a, q.argMax)
				res++
			}
		}

		delta = 0
		for i := m; i < n; i++ {
			q := st.Query(m, i)
			if prefixSums[i]+delta < prefixSums[m-1] {
				delta += -2 * q.min
				a[q.argMin] = -a[q.argMin]
				st.Update(a, q.argMin)
				res++
			}
		}

		fmt.Fprintln(out, res)
	}
}
