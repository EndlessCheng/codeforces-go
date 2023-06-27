package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1081/C
// https://codeforces.com/problemset/status/1081/problem/C
func TestCF1081C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 0
outputCopy
3
inputCopy
3 2 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1081C)
}
