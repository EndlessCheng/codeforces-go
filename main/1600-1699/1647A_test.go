package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1647/A
// https://codeforces.com/problemset/status/1647/problem/A
func TestCF1647A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1
2
3
4
5
outputCopy
1
2
21
121
212`
	testutil.AssertEqualCase(t, rawText, 0, CF1647A)
}
