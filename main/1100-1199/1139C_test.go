package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1139/C
// https://codeforces.com/problemset/status/1139/problem/C
func TestCF1139C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2 1
2 3 1
3 4 1
outputCopy
252
inputCopy
4 6
1 2 0
1 3 0
1 4 0
outputCopy
0
inputCopy
3 5
1 2 1
2 3 0
outputCopy
210`
	testutil.AssertEqualCase(t, rawText, 0, CF1139C)
}
