package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/300/A
// https://codeforces.com/problemset/status/300/problem/A
func TestCF300A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
-1 2 0
outputCopy
1 -1
1 2
1 0
inputCopy
4
-1 -2 -3 0
outputCopy
1 -1
2 -3 -2
1 0`
	testutil.AssertEqualCase(t, rawText, 0, CF300A)
}
