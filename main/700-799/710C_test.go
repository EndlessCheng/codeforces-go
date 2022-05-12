package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/710/C
// https://codeforces.com/problemset/status/710/problem/C
func TestCF710C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
1
inputCopy
3
outputCopy
2 1 4
3 5 7
6 9 8`
	testutil.AssertEqualCase(t, rawText, 0, CF710C)
}
