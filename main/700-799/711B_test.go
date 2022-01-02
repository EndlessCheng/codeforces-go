package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/711/B
// https://codeforces.com/problemset/status/711/problem/B
func TestCF711B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 0 2
3 5 7
8 1 6
outputCopy
9
inputCopy
4
1 1 1 1
1 1 0 1
1 1 1 1
1 1 1 1
outputCopy
1
inputCopy
4
1 1 1 1
1 1 0 1
1 1 2 1
1 1 1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 1, CF711B)
}
