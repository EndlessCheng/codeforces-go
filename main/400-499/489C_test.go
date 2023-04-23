package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/489/C
// https://codeforces.com/problemset/status/489/problem/C
func TestCF489C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 15
outputCopy
69 96
inputCopy
3 0
outputCopy
-1 -1
inputCopy
2 1
outputCopy
10 10`
	testutil.AssertEqualCase(t, rawText, 0, CF489C)
}
