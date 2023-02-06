package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1695/C
// https://codeforces.com/problemset/status/1695/problem/C
func TestCF1695C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
1
1 2
1 -1
1 4
1 -1 1 -1
3 4
1 -1 -1 -1
-1 1 1 -1
1 1 1 -1
3 4
1 -1 1 1
-1 1 -1 1
1 -1 1 1
outputCopy
NO
YES
YES
YES
NO
inputCopy
1
3 2
1 1
1 -1
-1 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF1695C)
}

func TestCompareCF1695C(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 9)
		m := rg.Int(1, 9)
		rg.NewLine()
		rg.IntMatrixInSet(n, m, []int{-1, 1})
		return rg.String()
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T int
		fmt.Fscan(in, &T)
		const N = 1010
		var grid [N][N]int
		var mx [N][N]int
		var mn [N][N]int
		for t := 0; t < T; t++ {
			var n, m int
			fmt.Fscan(in, &n, &m)

			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					fmt.Fscan(in, &grid[i][j])
				}
			}

			mx[0][0], mn[0][0] = grid[0][0], grid[0][0]
			for i := 1; i < n; i++ {
				mx[i][0] = mx[i-1][0] + grid[i][0]
				mn[i][0] = mx[i][0]
			}

			for j := 1; j < m; j++ {
				mx[0][j] = mx[0][j-1] + grid[0][j]
				mn[0][j] = mx[0][j]
			}

			for i := 1; i < n; i++ {
				for j := 1; j < m; j++ {
					mx[i][j] = max(mx[i-1][j], mx[i][j-1]) + grid[i][j]
					mn[i][j] = min(mn[i-1][j], mn[i][j-1]) + grid[i][j]
				}
			}

			if mx[n-1][m-1]%2 != 0 || mn[n-1][m-1] > 0 || mx[n-1][m-1] < 0 {
				fmt.Fprintln(out, "NO")
			} else {
				fmt.Fprintln(out, "YES")
			}
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1695C)
}
