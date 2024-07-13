package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1167/F
// https://codeforces.com/problemset/status/1167/problem/F
func TestCF1167F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 2 4 7
outputCopy
167
inputCopy
3
123456789 214365879 987654321
outputCopy
582491518`
	testutil.AssertEqualCase(t, rawText, 1, CF1167F)
}
