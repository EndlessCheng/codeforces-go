package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/519/D
// https://codeforces.com/problemset/status/519/problem/D
func TestCF519D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 -1 1 1 1 1 1 1 1 1 1 1 1 1 7 1 1 1 8 1 1 1 1 1 1
xabcab
outputCopy
2
inputCopy
1 1 -1 1 1 1 1 1 1 1 1 1 1 1 1 7 1 1 1 8 1 1 1 1 1 1
aaa
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 1, CF519D)
}
