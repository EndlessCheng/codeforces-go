package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/817/D
// https://codeforces.com/problemset/status/817/problem/D
func TestCF817D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 4 1
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF817D)
}
