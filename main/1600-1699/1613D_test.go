package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1613/D
// https://codeforces.com/problemset/status/1613/problem/D
func TestCF1613D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
0 2 1
2
1 0
5
0 0 0 0 0
4
0 1 2 3
outputCopy
4
2
31
7
inputCopy
1
7
1 3 2 3 0 0 1
outputCopy
9
inputCopy
1
6
0 3 2 2 1 3
outputCopy
7
inputCopy
1
9
0 1 0 1 1 3 0 3 1
outputCopy
133
inputCopy
1
3
0 1 1
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, -1, CF1613D)
}

func TestCompareCF1613D(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 3)
		rg.NewLine()
		rg.IntSlice(n, 0,n)
		return rg.String()
	}
	
	testutil.AssertEqualRunResultsInf(_t, inputGenerator, CF1613DAC, CF1613D)
}


func add13(a, b int) int {
	const MOD = 998244353
	return (a + b) % MOD
}

func CF1613DAC(in io.Reader, out io.Writer) {
	const SZ = 5e5 + 10
	var a [SZ]int
	var dp0 [SZ]int
	var dp1 [2][SZ]int
	var ntc int
	fmt.Fscan(in, &ntc)
	for t := 0; t < ntc; t++ {
		var n int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i <= n+3; i++ {
			dp0[i] = 0
			dp1[0][i] = 0
			dp1[1][i] = 0
		}

		var ans int = 0
		for i := n - 1; i >= 0; i-- {
			var x = a[i]
			// ways to continue a subsequence starting in a[i]=x
			var r0 int = 1             // if current mex is x-1
			var r1unrestricted int = 1 // if current mex is x+1 but x+2 is NOT in the sequence
			var r1restricted int = 1   // if current mex is x+1 and x+2 is in the sequence

			// next number is x-2 and next mex is x-1
			if x >= 2 {
				r0 = add13(r0, dp1[1][x-2])
			}
			// next number is x and next mex is x-1
			r0 = add13(r0, dp0[x])

			// next number is x and next mex x+1
			r1unrestricted = add13(r1unrestricted, dp1[0][x])
			// next number is x+1 and next mex x+2
			r1unrestricted = add13(r1unrestricted, dp1[0][x+1])
			// next number is x+2 and next mex is x+1
			r1unrestricted = add13(r1unrestricted, dp0[x+2])

			// next number is x and next mex is x+1
			r1restricted = add13(r1restricted, dp1[1][x])
			// next number is x+2 and next mex is x+1
			r1restricted = add13(r1restricted, dp0[x+2])

			if x == 1 {
				ans = add13(ans, r0)
			} else if x == 0 {
				ans = add13(ans, r1unrestricted)
			}

			dp0[x] = add13(dp0[x], r0)
			dp1[0][x] = add13(dp1[0][x], r1unrestricted)
			dp1[1][x] = add13(dp1[1][x], r1restricted)
		}
		fmt.Fprintf(out, "%d\n", ans)
	}
}
