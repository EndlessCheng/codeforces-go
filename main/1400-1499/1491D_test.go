package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1491/problem/D
// https://codeforces.com/problemset/status/1491/problem/D
func TestCF1491D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 4
3 6
1 6
6 2
5 5
outputCopy
YES
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1491D)
}
