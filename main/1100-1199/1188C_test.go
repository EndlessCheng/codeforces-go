package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1188/problem/C
// https://codeforces.com/problemset/status/1188/problem/C
func TestCF1188C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 7 3 5
outputCopy
8
inputCopy
5 5
1 10 100 1000 10000
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 1, CF1188C)
}
