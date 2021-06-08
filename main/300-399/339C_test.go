package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/339/C
// https://codeforces.com/problemset/status/339/problem/C
func TestCF339C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0000000101
3
outputCopy
YES
8 10 8
inputCopy
1000000000
2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF339C)
}
