package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math"
	"testing"
)

// https://codeforces.com/problemset/problem/1407/D
// https://codeforces.com/problemset/status/1407/problem/D
func TestCF1407D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 3 1 4 5
outputCopy
3
inputCopy
4
4 2 2 4
outputCopy
1
inputCopy
2
1 1
outputCopy
1
inputCopy
5
100 1 100 1 100
outputCopy
2
inputCopy
3
3 3 5
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF1407D)
}

func TestCompareCF1407D(t *testing.T) {
	return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 9)
		rg.NewLine()
		rg.IntSlice(n, 1, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		var u, d []int
		nxt := make([][]int, n)

		upd := func(i int, add func(int, int)) {
			for len(u) > 0 && a[u[len(u)-1]] < a[i] {
				u = u[:len(u)-1]
			}
			for len(d) > 0 && a[d[len(d)-1]] > a[i] {
				d = d[:len(d)-1]
			}
			if len(u) > 0 {
				add(u[len(u)-1], i)
			}
			if len(d) > 0 {
				add(d[len(d)-1], i)
			}
			u, d = append(u, i), append(d, i)
		}

		for i := 0; i < n; i++ {
			upd(i, func(i, j int) { nxt[i] = append(nxt[i], j) })
		}
		for i := n - 1; i >= 0; i-- {
			upd(i, func(i, j int) { nxt[j] = append(nxt[j], i) })
		}

		dp := make([]int, n)
		for i := n - 2; i >= 0; i-- {
			dp[i] = math.MaxInt32
			for _, j := range nxt[i] {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}

		fmt.Fprintln(out, dp[0])
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1407D)
}
