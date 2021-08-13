package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1036/F
// https://codeforces.com/problemset/status/1036/problem/F
func TestCF1036F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
2
72
10
outputCopy
2
1
61
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1036F)
}
