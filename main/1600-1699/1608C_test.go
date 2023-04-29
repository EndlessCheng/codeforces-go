package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/1608/C
// https://codeforces.com/problemset/status/1608/problem/C
func TestCF1608C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 2 3 4
1 2 3 4
4
11 12 20 21
44 22 11 30
1
1000000000
1000000000
outputCopy
0001
1111
1
inputCopy
1
3
1 2 3
2 1 3
outputCopy
001
inputCopy
1
6
1 2 3 4 5 6
2 4 1 6 3 5
outputCopy
111111`
	testutil.AssertEqualCase(t, rawText, 0, CF1608C)
}

func TestCompareCF1608C(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 9)
		rg.NewLine()
		rg.IntSliceOrdered(n, 1, n, true, true)
		rg.UniqueSlice(n, 1, n)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, solveCF1608C, CF1608C)
}

func solveCF1608C(input io.Reader, output io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	type pair struct{ a, b int }
	var n int
	fmt.Fscan(input, &n, &n)
	a := make([]pair, n)
	aa := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &a[i].a)
		aa[i] = a[i].a
		a[i].b = i
	}

	b := make([]pair, n)
	bb := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &b[i].a)
		bb[i] = b[i].a
		b[i].b = i
	}

	sort.Slice(a, func(i, j int) bool { return a[i].a < a[j].a })
	sort.Slice(b, func(i, j int) bool { return b[i].a < b[j].a })

	ans := make([]int, n)
	i := n - 1
	mina := min(a[i].a, aa[b[i].b])
	minb := min(b[i].a, bb[a[i].b])

	for i >= 0 {
		if a[i].a >= mina || bb[a[i].b] > minb {
			ans[a[i].b] = 1
			mina = min(mina, a[i].a)
			minb = min(minb, bb[a[i].b])
		}
		if b[i].a >= minb || aa[b[i].b] > mina {
			ans[b[i].b] = 1
			minb = min(minb, b[i].a)
			mina = min(mina, aa[b[i].b])
		}
		i--
	}

	for _, v := range ans {
		fmt.Fprint(output, v)
	}
	fmt.Fprintln(output)
}
