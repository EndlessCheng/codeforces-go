package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1804/D
// https://codeforces.com/problemset/status/1804/problem/D
func TestCF1804D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
0100
1100
0110
1010
1011
outputCopy
7 10
inputCopy
1 8
01011100
outputCopy
3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1804D)
}

func TestCompareCF1804D(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.Int(1, 1)
		m := rg.Int(8, 8) 
		rg.NewLine()
		rg.Str(m, m, '0', '1')
		return rg.String()
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, m, mini, maxi int
		var s string
		Fscan(in, &n, &m, &s)
		k := strings.Count(s, "1")
		c := 0
		for i := 0; i+1 < m; i++ {
			if s[i] == '1' && s[i+1] == '1' {
				c++
				i++
			}
		}
		mini += k - min(c, m/4)
		c = 0
		for i := 0; i+1 < m; i++ {
			if s[i] != '1' || s[i+1] != '1' {
				c++
				i++
			}
		}
		maxi += k - max(0, m/4-c)
		Fprintln(out, mini,maxi)
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1804D)
}
