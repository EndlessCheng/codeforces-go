package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1251/problem/C
// https://codeforces.com/problemset/status/1251/problem/C
func TestCF1251C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0709
1337
246432
outputCopy
0079
1337
234642`
	testutil.AssertEqualCase(t, rawText, 0, CF1251C)
}
