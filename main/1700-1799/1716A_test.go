package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1716/problem/A
// https://codeforces.com/problemset/status/1716/problem/A
func TestCF1716A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
3
4
12
outputCopy
2
1
2
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1716A)
}
