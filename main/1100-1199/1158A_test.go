package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1158/A
// https://codeforces.com/problemset/status/1158/problem/A
func TestCF1158A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2 1
3 4
outputCopy
12
inputCopy
2 2
0 1
1 0
outputCopy
-1
inputCopy
2 3
1 0
1 1 2
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1158A)
}
