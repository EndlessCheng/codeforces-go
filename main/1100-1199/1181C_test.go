package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// http://codeforces.com/problemset/problem/1181/C
// https://codeforces.com/problemset/status/1181/problem/C
func TestCF1181C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
aaa
bbb
ccb
ddd
outputCopy
6
inputCopy
6 1
a
a
b
b
c
c
outputCopy
1
inputCopy
6 5
bbaaa
baaab
bbbab
aabbb
aaaaa
aaaaa
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1181C)
}

func TestCompareCF1181C(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		m := rg.Int(1, 9)
		rg.NewLine()
		for i := 0; i < n; i++ {
			rg.Str(m, m, 'a', 'c')
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, m, ans int
		Fscan(in, &n, &m)
		s := make([]string, n+1)
		for i := 1; i <= n; i++ {

			Fscan(in, &s[i])
			s[i] = " " + s[i]
		}
		a := [100][100]int{}
		f := [100][100]int{}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				if i == 1 || s[i][j] != s[i-1][j] {
					a[i][j] = 1
				} else {
					a[i][j] = a[i-1][j] + 1
				}
			}

		}

		for i := 1; i <= n; i++ {
			la := 0
			for j := 1; j <= m; j++ {
				t := a[i][j]
				if a[i-t][j] != t || a[i-2*t][j] < t {
					la = 0
					continue
				}
				f[i][j] = 1
				if la == t && s[i][j-1] == s[i][j] && s[i-t][j-1] == s[i-t][j] && s[i-2*t][j-1] == s[i-2*t][j] {
					f[i][j] += f[i][j-1]

				}
				la = t
				ans += f[i][j]
			}
		}
		Fprint(out, ans)
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1181C)
}
