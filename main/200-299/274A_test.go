package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/274/A
// https://codeforces.com/problemset/status/274/problem/A
func TestCF274A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2
2 3 6 5 4 10
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF274A)
}
