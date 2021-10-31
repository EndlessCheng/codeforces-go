package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1005/C
// https://codeforces.com/problemset/status/1005/problem/C
func TestCF1005C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 7 1 5 4 9
outputCopy
1
inputCopy
5
1 2 3 4 5
outputCopy
2
inputCopy
1
16
outputCopy
1
inputCopy
4
1 1 1 1023
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1005C)
}
