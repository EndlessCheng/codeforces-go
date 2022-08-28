package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/6/A
// https://codeforces.com/problemset/status/6/problem/A
func TestCF6A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2 1 3
outputCopy
TRIANGLE
inputCopy
7 2 2 4
outputCopy
SEGMENT
inputCopy
3 5 9 1
outputCopy
IMPOSSIBLE`
	testutil.AssertEqualCase(t, rawText, 0, CF6A)
}
