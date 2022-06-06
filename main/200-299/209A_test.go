package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/209/A
// https://codeforces.com/problemset/status/209/problem/A
func TestCF209A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
6
inputCopy
4
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF209A)
}
