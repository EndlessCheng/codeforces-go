package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1332/E
// https://codeforces.com/problemset/status/1332/problem/E
func TestCF1332E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 1 1
outputCopy
1
inputCopy
1 2 1 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1332E)
}
