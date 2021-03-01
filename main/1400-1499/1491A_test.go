package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1491/problem/A
// https://codeforces.com/problemset/status/1491/problem/A
func TestCF1491A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 1 0 1 0
2 3
1 2
2 3
2 1
2 5
outputCopy
1
0
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1491A)
}
