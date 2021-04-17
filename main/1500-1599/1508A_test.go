package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1508/A
// https://codeforces.com/problemset/status/1508/problem/A
func TestCF1508A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1
00
11
01
3
011001
111010
010001
outputCopy
010
011001010`
	testutil.AssertEqualCase(t, rawText, 0, CF1508A)
}
