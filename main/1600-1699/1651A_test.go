package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1651/problem/A
// https://codeforces.com/problemset/status/1651/problem/A
func TestCF1651A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
1
outputCopy
7
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1651A)
}
