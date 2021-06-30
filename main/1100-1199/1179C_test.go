package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1179/C
// https://codeforces.com/problemset/status/1179/problem/C
func TestCF1179C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
1
1
1
1 1 100
outputCopy
100
inputCopy
1 1
1
1
1
2 1 100
outputCopy
-1
inputCopy
4 6
1 8 2 4
3 3 6 1 5 2
3
1 1 1
2 5 10
1 1 6
outputCopy
8
-1
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1179C)
}
