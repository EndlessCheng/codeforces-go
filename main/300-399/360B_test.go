package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/360/B
// https://codeforces.com/problemset/status/360/problem/B
func TestCF360B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
4 7 4 7 4
outputCopy
0
inputCopy
3 1
-100 0 100
outputCopy
100
inputCopy
6 3
1 2 3 7 8 9
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF360B)
}
