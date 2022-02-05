package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/contest/1622/problem/D
// https://codeforces.com/problemset/status/1622/problem/D
func TestCF1622D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
1100110
outputCopy
16
inputCopy
5 0
10010
outputCopy
1
inputCopy
8 1
10001000
outputCopy
10
inputCopy
10 8
0010011000
outputCopy
1
inputCopy
5 2
11101
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, -1, CF1622D)
}

func TestCompareCF1622D(t *testing.T) {
	return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 20)
		rg.Int(0, n)
		rg.NewLine()
		rg.Str(n, n, '0', '1')
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		const MOD = 998244353
		var n, k int
		var s string
		Fscan(in, &n, &k, &s)
		if k == 0 {
			Fprint(out, 1)
			return
		}

		C := make([][]int64, n+1)
		for i := 0; i <= n; i++ {
			C[i] = make([]int64, k+1)
			for j := 0; j <= k; j++ {
				if j == 0 {
					C[i][j] = 1
				} else if i == 0 {
					C[i][j] = 0
				} else {
					C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
				}
			}
		}

		L := make([]int64, n)
		cur := 0
		sum1 := 0
		for i := 0; i < n; i++ {
			sum1 += int(s[i] - '0')
			for sum1 > k {
				sum1 -= int(s[cur] - '0')
				cur += 1
			}
			L[i] = int64(i - cur + 1)
		}

		dp := make([]int64, n)
		sum1 = 0
		for i := 0; i < n; i++ {
			sum1 += int(s[i] - '0')
			if sum1 < k {
				dp[i] = 1
			} else if sum1 == k {
				dp[i] = C[L[i]][k]
			} else {
				if s[i] == '0' {
					dp[i] = (dp[i-1] + C[L[i]-1][k-1]) % MOD
				} else {
					dp[i] = (dp[i-1] + C[L[i]-1][k]) % MOD
				}
			}
		}
		Fprint(out, dp[n-1])
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1622D)
}
