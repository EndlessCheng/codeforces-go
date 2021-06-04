package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1469/C
// https://codeforces.com/problemset/status/1469/problem/C
func TestCF1469C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6 3
0 0 2 5 1 1
2 3
0 2
3 2
3 0 2
outputCopy
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1469C)
}

func TestCompareCF1469C(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(2, 7)
		rg.Int(2,5)
		rg.NewLine()
		rg.IntSlice(n, 0, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T, n, k int
		for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &n,&k)
			h := make([]int, n)
			for i := range h {
				Fscan(in, &h[i])
			}
			var f func(p ,v int) bool
			f = func(p ,v int) bool {
				if p == n-1 {
					return v==h[n-1]
				}
				if v < h[p] || v >= h[p]+k {
					return false
				}
				for i := v-(k-1);i< v+k;i++ {
					if f(p+1,i) {
						return true
					}
				}
				return false
			}
			if f(0, h[0]) {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}

		}
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//3
//6 3
//0 0 2 5 1 1
//2 3
//0 2
//3 2
//3 0 2
//outputCopy
//YES
//YES
//NO`
//	testutil.AssertEqualCase(t, rawText, 0, CF1469C)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1469C)
}
