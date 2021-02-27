package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1415/D
// https://codeforces.com/problemset/status/1415/problem/D
func TestCF1415D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 5 6 8
outputCopy
1
inputCopy
3
1 2 3
outputCopy
-1
inputCopy
5
1 2 4 6 20
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1415D)
}
