package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/463/C
// https://codeforces.com/problemset/status/463/problem/C
func TestCF463C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1 1
2 1 1 0
1 1 1 0
1 0 0 1
outputCopy
12
2 2 3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF463C)
}
