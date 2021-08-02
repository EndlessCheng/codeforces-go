package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/908/D
// https://codeforces.com/problemset/status/908/problem/D
func TestCF908D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1
outputCopy
2
inputCopy
3 1 4
outputCopy
370000006`
	testutil.AssertEqualCase(t, rawText, 0, CF908D)
}
