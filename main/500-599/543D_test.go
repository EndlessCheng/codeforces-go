package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/543/D
// https://codeforces.com/problemset/status/543/problem/D
func TestCF543D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1
outputCopy
4 3 3
inputCopy
5
1 2 3 4
outputCopy
5 8 9 8 5
inputCopy
70
1 2 2 4 4 6 6 8 9 9 11 11 13 13 15 15 17 17 19 19 21 22 22 24 24 26 27 27 29 29 31 31 33 34 34 36 37 37 39 39 41 42 42 44 44 46 47 47 49 50 50 52 52 54 54 56 57 57 59 60 60 62 63 63 65 65 67 68 68
outputCopy
0 1000000005 0 499999996 249999999 749999986 374999994 874999963 999999938 499999970 62499881 531249945 93749781 546874895 109374581 554687295 117186681 558593345 121092131 560546070 123043656 124995179 562497594 125968539 562984274 126450416 126932291 563466150 127163621 563581815 127260071 563630040 127269866 127279659 563639834 127207694 127135727 563567868 126946019 563473014 126543716 126141411 563070710 125325359 562662684 123687534 122049707 561024858 118771194 115492679 557746344 108934221 55446711...
inputCopy
4
1 2 1
outputCopy
`
	testutil.AssertEqualCase(t, rawText, -1, CF543D)
}

//func TestCompareCF543D(_t *testing.T) {
//	//return
//	testutil.DebugTLE = 0
//
//	inputGenerator := func() string {
//		//return ``
//		rg := testutil.NewRandGenerator()
//		n := rg.Int(1, 4)
//		rg.NewLine()
//		for i := 2; i <= n; i++ {
//			rg.Int(1, i-1)
//		}
//		return rg.String()
//	}
//
//	testutil.AssertEqualRunResultsInf(_t, inputGenerator, CF543DAC, CF543D)
//}
//
//func CF543DAC(_r io.Reader, _w io.Writer) {
//	in := bufio.NewReader(_r)
//	out := bufio.NewWriter(_w)
//	defer out.Flush()
//
//	var n, v int
//	Fscan(in, &n)
//	g := make([][]int, n)
//	for w := 1; w < n; w++ {
//		Fscan(in, &v)
//		v--
//		g[v] = append(g[v], w)
//	}
//
//	dp := make([]int64, n)
//	ex := make([]int64, n)
//	var f func(int)
//	f = func(v int) {
//		z := false
//		dp[v] = 1
//		ex[v] = 1
//		for _, w := range g[v] {
//			f(w)
//			dw := dp[w] + 1
//			dp[v] = dp[v] * dw % mod
//			if z || dw != mod {
//				ex[v] = ex[v] * dw % mod
//			} else {
//				z = true
//			}
//		}
//	}
//	f(0)
//
//	ans := make([]int64, n)
//	var reroot func(int, int64)
//	reroot = func(v int, dpFa int64) {
//		ans[v] = dp[v] * (dpFa + 1) % mod
//		for _, w := range g[v] {
//			df := int64(0)
//			if dp[w]+1 == mod {
//				df = ex[v] * (dpFa + 1) % mod
//			} else {
//				df = ans[v] * pow(dp[w]+1, mod-2) % mod
//			}
//			reroot(w, df)
//		}
//	}
//	reroot(0, 0)
//	for _, v := range ans {
//		Fprint(out, v, " ")
//	}
//}
