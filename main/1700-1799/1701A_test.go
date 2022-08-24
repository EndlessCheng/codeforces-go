package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1701/problem/A
// https://codeforces.com/problemset/status/1701/problem/A
func TestCF1701A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 0
0 0
1 0
0 1
1 1
1 1
outputCopy
0
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1701A)
}
