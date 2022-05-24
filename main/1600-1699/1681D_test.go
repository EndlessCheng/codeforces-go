package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1681/problem/D
// https://codeforces.com/problemset/status/1681/problem/D
func TestCF1681D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1
outputCopy
-1
inputCopy
3 2
outputCopy
4
inputCopy
13 42
outputCopy
12
inputCopy
17 992
outputCopy
14`
	testutil.AssertEqualCase(t, rawText, -1, CF1681D)
}
