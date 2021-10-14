package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1558/B
// https://codeforces.com/problemset/status/1558/problem/B
func TestCF1558B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 998244353
outputCopy
5
inputCopy
5 998244353
outputCopy
25
inputCopy
42 998244353
outputCopy
793019428
inputCopy
787788 100000007
outputCopy
94810539`
	testutil.AssertEqualCase(t, rawText, 0, CF1558B)
}
