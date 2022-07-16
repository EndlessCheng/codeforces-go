package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/contest/1705/problem/E
// https://codeforces.com/problemset/status/1705/problem/E
func TestCF1705E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
2 2 2 4 5
2 3
5 3
4 1
1 4
outputCopy
6
5
4
5
inputCopy
2 1
200000 1
2 200000
outputCopy
200001
inputCopy
1 5
2
1 5
1 5
1 1
1 2
1 2
outputCopy
5
5
1
2
2
inputCopy
10 3
5 5 5 5 5 5 4 2 5 5
10 4
2 3
1 3
outputCopy
8
7
7
inputCopy
5 2
3 10 9 15 12
1 8
2 6
outputCopy
15
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1705E)
}

func TestCompareCF1705E(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		q := rg.Int(1,5)
		rg.NewLine()
		rg.IntSlice(n, 1, 20)
		for i := 0; i < q; i++ {
			rg.Int(1,n)
			rg.Int(1,20)
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n,q int
		Fscan(in, &n, &q)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		cnt := [100]int{}
		for _, v := range a[1:] {
			cnt[v]++
		}
		for i := 0; i < q; i++ {
			var i,v int
			Fscan(in, &i, &v)
			cnt[a[i]]--
			a[i] = v
			cnt[v]++
			c := cnt
			ans := 1
			for i := 2; i < 100; i++ {
				c[i] += c[i-1]/2
				if c[i] > 0 {
					ans = i
				}
			}
			Fprintln(out, ans)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1705E)
}
