package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1081/problem/E
// https://codeforces.com/problemset/status/1081/problem/E
func TestCF1081E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 11 44
outputCopy
Yes
4 5 16 11 64 44
inputCopy
2
9900
outputCopy
Yes
100 9900
inputCopy
6
314 1592 6535
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1081E)
}
