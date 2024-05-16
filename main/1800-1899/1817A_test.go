package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/contest/1817/problem/A
// https://codeforces.com/problemset/status/1817/problem/A
func TestCF1817A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 8
1 2 4 3 3 5 6 2 1
1 3
1 4
2 5
6 6
3 7
7 8
1 8
8 8
outputCopy
3
4
3
1
4
2
7
1
inputCopy
9 1
1 2 4 3 3 5 6 2 1
8 8
outputCopy
1
inputCopy
3 1
3 3 4
1 3
outputCopy
3
inputCopy
4 1
2 4 4 5
1 4
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, -1, CF1817A)
}

func TestCompare_CF1817A(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 4)
		rg.Int(1,1)
		rg.NewLine()
		rg.IntSlice(n, 1, 5)
		l := rg.Int(1,n)
		rg.Int(l,n)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, q, l, r int
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for qi := 0; qi < q; qi++ {
			Fscan(in, &l, &r)
			b := append([]int{-2, -1}, a[l-1 : r]...)
			ans := 0
			for i := 2; i < len(b); i++ {
				if b[i] <= b[i-1] && b[i-1] <= b[i-2] {
					continue
				}
				ans++
			}
			Fprintln(out, ans)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1817A)
}

