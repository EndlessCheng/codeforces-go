package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1017/C
// https://codeforces.com/problemset/status/1017/problem/C
func TestCF1017C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
3 4 1 2
inputCopy
2
outputCopy
2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1017C)
}
