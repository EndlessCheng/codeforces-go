package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1283/problem/A
// https://codeforces.com/problemset/status/1283/problem/A
func TestCF1283A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
23 55
23 0
0 1
4 20
23 59
outputCopy
5
60
1439
1180
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1283A)
}
