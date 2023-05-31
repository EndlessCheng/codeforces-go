package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/190/D
// https://codeforces.com/problemset/status/190/problem/D
func TestCF190D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2 1 2
outputCopy
3
inputCopy
5 3
1 2 1 1 3
outputCopy
2
inputCopy
3 1
1 1 1
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF190D)
}
