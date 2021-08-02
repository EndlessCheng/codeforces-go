package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1070/A
// https://codeforces.com/problemset/status/1070/problem/A
func TestCF1070A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
13 50
outputCopy
699998
inputCopy
61 2
outputCopy
1000000000000000000000000000001
inputCopy
15 50
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1070A)
}
