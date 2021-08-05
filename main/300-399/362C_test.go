package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/362/C
// https://codeforces.com/problemset/status/362/problem/C
func TestCF362C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 0 3 1 2
outputCopy
3 2
inputCopy
5
1 2 3 4 0
outputCopy
3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF362C)
}
