package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1401/C
// https://codeforces.com/problemset/status/1401/problem/C
func TestCF1401C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
8
6
4 3 6 6 2 9
4
4 5 6 7
5
7 5 2 2 4
outputCopy
YES
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1401C)
}
