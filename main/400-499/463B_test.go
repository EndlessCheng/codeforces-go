package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/463/B
// https://codeforces.com/problemset/status/463/problem/B
func TestCF463B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 4 3 2 4
outputCopy
4
inputCopy
3
4 4 4
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF463B)
}
