package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/463/D
// https://codeforces.com/problemset/status/463/problem/D
func TestCF463D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 4 2 3
4 1 2 3
1 2 4 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF463D)
}
