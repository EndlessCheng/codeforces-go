package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/605/A
// https://codeforces.com/problemset/status/605/problem/A
func TestCF605A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 1 2 5 3
outputCopy
2
inputCopy
4
4 1 3 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF605A)
}
