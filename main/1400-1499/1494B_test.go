package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1494/B
// https://codeforces.com/problemset/status/1494/problem/B
func TestCF1494B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 2 5 3 1
3 0 0 0 0
4 4 1 4 0
2 1 1 1 1
outputCopy
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1494B)
}
