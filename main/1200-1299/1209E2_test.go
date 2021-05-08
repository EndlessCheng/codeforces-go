package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

// https://codeforces.com/contest/1209/problem/E2
// https://codeforces.com/problemset/status/1209/problem/E2
func TestCF1209E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3
2 5 7
4 2 4
3 6
4 1 5 2 10 4
8 6 6 4 9 10
5 4 9 5 8 7
3 3
9 9 9
1 1 1
1 1 1
outputCopy
12
29
27
inputCopy
1
4 2
4 1
1 2
4 3
2 5
outputCopy
15
`
	testutil.AssertEqualCase(t, rawText, -1, CF1209E2)
}

func TestCompareCF1209E2(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 4)
		m := rg.Int(1, 5)
		rg.NewLine()
		rg.IntMatrix(n, m, 1, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		solve := func(Case int) {
			var n, m int
			Fscan(in, &n, &m)
			a := make([][]int, n)
			for i := range a {
				a[i] = make([]int, m)
				for j := range a[i] {
					Fscan(in, &a[i][j])
				}
			}
			mxall := 0
			var f func(col int)
			f = func(col int) {
				if col == m {
					s := 0
					for _, r := range a {
						mx := int(-1e9)
						for _, v := range r {
							if v > mx {
								mx = v
							}
						}
						s += mx
					}
					if s > mxall {
						mxall = s
					}
					return
				}

				for k := 0; k < n; k++ {
					f(col + 1)
					tmp := a[0][col]
					for i := 1; i < n; i++ {
						a[i-1][col] = a[i][col]
					}
					a[n-1][col] = tmp
				}

			}
			f(0)
			Fprintln(out, mxall)
		}

		T := 1
		Fscan(in, &T)
		for Case := 1; Case <= T; Case++ {
			solve(Case)
		}

		_leftData, _ := ioutil.ReadAll(in)
		if _s := strings.TrimSpace(string(_leftData)); _s != "" {
			panic("有未读入的数据：\n" + _s)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1209E2)
}
