package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/846/A
// https://codeforces.com/problemset/status/846/problem/A
func TestCF846A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 0 1
outputCopy
3
inputCopy
6
0 1 0 0 1 0
outputCopy
4
inputCopy
1
0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF846A)
}
