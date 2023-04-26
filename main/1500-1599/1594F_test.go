package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1594/F
// https://codeforces.com/problemset/status/1594/problem/F
func TestCF1594F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1
1 1 2
100 50 200
56220 47258 14497
outputCopy
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1594F)
}

func TestCompareCF1594F(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		s := rg.Int(1, 3)
		rg.Int(1, s)
		rg.Int(1, 3)
		rg.NewLine()
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T int
		var s, n, k int64
		for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &s, &n, &k)
			if s==k||s<n/k*k+n {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确


	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1594F)

	// for hacking, write the hacked codes in runBF
	//testutil.AssertEqualRunResultsInf(_t, inputGenerator, run, runBF)
}
