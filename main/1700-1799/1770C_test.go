package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1770/problem/C
// https://codeforces.com/problemset/status/1770/problem/C
func TestCF1770C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
5 7 10
3
3 3 4
outputCopy
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1770C)
}
