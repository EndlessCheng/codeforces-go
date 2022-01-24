package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/715/A
// https://codeforces.com/problemset/status/715/problem/A
func TestCF715A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
14
16
46
inputCopy
2
outputCopy
999999999999999998
44500000000
inputCopy
4
outputCopy
2
17
46
97`
	testutil.AssertEqualCase(t, rawText, 0, CF715A)
}
