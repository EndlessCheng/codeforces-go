package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/1467/C
// https://codeforces.com/problemset/status/1467/problem/C
func TestCF1467C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4 1
1 2
6 3 4 5
5
outputCopy
20
inputCopy
3 2 2
7 5 4
2 9
7 1
outputCopy
29
inputCopy
1 1 2
2
2
1 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF1467C)
}

func TestCompareCF1467C(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n1,n2,n3 := rg.Int(1,2), rg.Int(1,2), rg.Int(1,2)
		rg.NewLine()
		rg.IntSlice(n1,1,5)
		rg.IntSlice(n2,1,5)
		rg.IntSlice(n3,1,5)
		return rg.String()
	}

	// 暴力算法
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	max64 := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	runBF := func(in io.Reader, out io.Writer) {

		var n1, n2, n3 int
		var sum, sumA, sumB, sumC int64
		fmt.Fscan(in, &n1, &n2, &n3)

		a := make([]int, n1)
		b := make([]int, n2)
		c := make([]int, n3)
		for i := 0; i < n1; i++ {
			fmt.Fscan(in, &a[i])
			sumA += int64(a[i])
		}
		for i := 0; i < n2; i++ {
			fmt.Fscan(in, &b[i])
			sumB += int64(b[i])
		}
		for i := 0; i < n3; i++ {
			fmt.Fscan(in, &c[i])
			sumC += int64(c[i])
		}
		sort.Ints(a)
		sort.Ints(b)
		sort.Ints(c)

		sum = sumA + sumB + sumC - 2*int64(a[0]+b[0]+c[0]-max(max(a[0], b[0]), c[0]))
		sum = max64(sum, sumA+sumB-sumC)
		sum = max64(sum, sumB+sumC-sumA)
		sum = max64(sum, sumC+sumA-sumB)
		fmt.Fprintln(out, sum)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1467C)
}
