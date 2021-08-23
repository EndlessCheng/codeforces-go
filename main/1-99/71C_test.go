package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/71/C
// https://codeforces.com/problemset/status/71/problem/C
func TestCF71C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1 1
outputCopy
YES
inputCopy
6
1 0 1 1 1 0
outputCopy
YES
inputCopy
6
1 0 0 1 0 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF71C)
}
