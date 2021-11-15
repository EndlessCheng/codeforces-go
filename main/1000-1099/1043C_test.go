package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1043/C
// https://codeforces.com/problemset/status/1043/problem/C
func TestCF1043C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
bbab
outputCopy
0 1 1 0
inputCopy
aaaaa
outputCopy
1 0 0 0 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1043C)
}
