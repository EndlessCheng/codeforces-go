package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1062/B
// https://codeforces.com/problemset/status/1062/problem/B
func TestCF1062B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
20
outputCopy
10 2
inputCopy
5184
outputCopy
6 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1062B)
}
