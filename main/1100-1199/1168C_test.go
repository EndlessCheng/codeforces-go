package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1168/problem/C
// https://codeforces.com/problemset/status/1168/problem/C
func TestCF1168C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 3 0 2 1
1 3
2 4
1 4
outputCopy
Fou
Shi
Shi`
	testutil.AssertEqualCase(t, rawText, 0, CF1168C)
}
