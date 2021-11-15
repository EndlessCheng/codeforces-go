package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/891/A
// https://codeforces.com/problemset/status/891/problem/A
func TestCF891A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 2 3 4 6
outputCopy
5
inputCopy
4
2 4 6 8
outputCopy
-1
inputCopy
3
2 6 9
outputCopy
4
inputCopy
4
1 1 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF891A)
}
