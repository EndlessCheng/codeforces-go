package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/45/G
// https://codeforces.com/problemset/status/45/problem/G
func TestCF45G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
outputCopy
1 2 2 1 1 1 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF45G)
}
