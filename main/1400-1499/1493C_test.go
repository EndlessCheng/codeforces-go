package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1493/C
// https://codeforces.com/problemset/status/1493/problem/C
func TestCF1493C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 2
abcd
3 1
abc
4 3
aaaa
9 3
abaabaaaa
outputCopy
acac
abc
-1
abaabaaab`
	testutil.AssertEqualCase(t, rawText, 0, CF1493C)
}

func TestCompareCF1493C(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 6)
		rg.Int(1,n)
		rg.NewLine()
		rg.Str(n,n,'a','z')
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var T, n, k int
		var s []byte

		for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &n, &k, &s)
			if n%k > 0 {
				Fprintln(out, -1)
				continue
			}
			Fprintln(out, "")
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1493C)
}
