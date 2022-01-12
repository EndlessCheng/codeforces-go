package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1088/C
// https://codeforces.com/problemset/status/1088/problem/C
func TestCF1088C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
outputCopy
0
inputCopy
3
7 6 3
outputCopy
2
1 1 1
2 2 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1088C)
}
