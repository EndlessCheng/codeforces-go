package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1516/C
// https://codeforces.com/problemset/status/1516/problem/C
func TestCF1516C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6 3 9 12
outputCopy
1
2
inputCopy
2
1 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1516C)
}
