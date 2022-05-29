package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1310/A
// https://codeforces.com/problemset/status/1310/problem/A
func TestCF1310A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 7 9 7 8
5 2 5 7 5
outputCopy
6
inputCopy
5
1 2 3 4 5
1 1 1 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1310A)
}
