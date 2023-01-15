package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1691/D
// https://codeforces.com/problemset/status/1691/problem/D
func TestCF1691D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
-1 1 -1 2
5
-1 2 -3 2 -1
3
2 3 -1
outputCopy
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1691D)
}
