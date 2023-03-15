package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1054/D
// https://codeforces.com/problemset/status/1054/problem/D
func TestCF1054D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 3 0
outputCopy
5
inputCopy
6 3
1 4 4 7 3 4
outputCopy
19`
	testutil.AssertEqualCase(t, rawText, 0, CF1054D)
}
