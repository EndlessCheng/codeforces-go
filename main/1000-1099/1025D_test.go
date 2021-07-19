package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1025/D
// https://codeforces.com/problemset/status/1025/problem/D
func TestCF1025D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 6 9 18 36 108
outputCopy
Yes
inputCopy
2
7 17
outputCopy
No
inputCopy
9
4 8 10 12 15 18 33 44 81
outputCopy
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1025D)
}
