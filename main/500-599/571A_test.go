package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/571/A
// https://codeforces.com/problemset/status/571/problem/A
func TestCF571A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1 2
outputCopy
4
inputCopy
1 2 3 1
outputCopy
2
inputCopy
10 2 1 7
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF571A)
}
