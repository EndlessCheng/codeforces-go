package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1610/D
// https://codeforces.com/problemset/status/1610/problem/D
func TestCF1610D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2 4 7
outputCopy
10
inputCopy
10
12391240 103904 1000000000 4142834 12039 142035823 1032840 49932183 230194823 984293123
outputCopy
996`
	testutil.AssertEqualCase(t, rawText, 0, CF1610D)
}
