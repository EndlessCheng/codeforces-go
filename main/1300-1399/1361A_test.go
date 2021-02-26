package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1361/A
// https://codeforces.com/problemset/status/1361/problem/A
func TestCF1361A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2
2 3
3 1
2 1 3
outputCopy
2 1 3
inputCopy
3 3
1 2
2 3
3 1
1 1 1
outputCopy
-1
inputCopy
5 3
1 2
2 3
4 5
2 1 2 2 1
outputCopy
2 5 1 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1361A)
}
