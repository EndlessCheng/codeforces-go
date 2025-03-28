// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/2005/D
// https://codeforces.com/problemset/status/2005/problem/D?friends=on
func Test_cf2005D(t *testing.T) {
	testCases := [][2]string{
		{
			`5
8
11 4 16 17 3 24 25 8
8 10 4 21 17 18 25 21
4
6 4 24 13
15 3 1 14
2
13 14
5 8
8
20 17 15 11 21 10 3 7
9 9 4 20 14 9 13 1
2
18 13
15 20`,
			`2 36
3 2
2 3
2 36
6 1`,
		},
		{
			`1
2
5 5
4 4`,
			`9 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2005D)
}

func TestCompare_cf2005D(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.One()
		n := rg.Int(1, 9)
		rg.NewLine()
		rg.IntSlice(n, 1, 9)
		rg.IntSlice(n, 1, 9)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		gcd := func(a, b int) int {
			for a != 0 {
				a, b = b%a, a
			}
			return b
		}
		solve := func(Case int) {
			var n int
			Fscan(in, &n)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}
			b := make([]int, n)
			for i := range b {
				Fscan(in, &b[i])
			}
			f := func(a []int) int {
				g := 0
				for _, v := range a {
					g = gcd(g, v)
				}
				return g
			}
			ans,cnt := 0,0
			for l := 0; l < n; l++ {
				for r := l; r < n; r++ {
					for i := l; i <= r; i++ {
						a[i], b[i] = b[i], a[i]
					}
					
					res:=f(a) + f(b)
					if res > ans {
						ans=res
						cnt = 1
					}else if res == ans {
						cnt++
					}

					for i := l; i <= r; i++ {
						a[i], b[i] = b[i], a[i]
					}
				}
			}
			Fprintln(out, ans,cnt)
		}

		T := 1
		Fscan(in, &T)
		for Case := 1; Case <= T; Case++ {
			solve(Case)
		}

		_leftData, _ := io.ReadAll(in)
		if _s := strings.TrimSpace(string(_leftData)); _s != "" {
			panic("有未读入的数据：\n" + _s)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf2005D)
}
