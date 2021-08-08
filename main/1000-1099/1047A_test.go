package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1047/A
// https://codeforces.com/problemset/status/1047/problem/A
func TestCF1047A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
1 1 1
inputCopy
233
outputCopy
77 77 79`
	testutil.AssertEqualCase(t, rawText, 0, CF1047A)
}
