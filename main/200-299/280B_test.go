package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/280/B
// https://codeforces.com/problemset/status/280/problem/B
func TestCF280B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 2 1 4 3
outputCopy
7
inputCopy
5
9 8 3 5 7
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF280B)
}
