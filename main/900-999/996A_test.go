package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/996/A
// https://codeforces.com/problemset/status/996/problem/A
func TestCF996A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
125
outputCopy
3
inputCopy
43
outputCopy
5
inputCopy
1000000000
outputCopy
10000000`
	testutil.AssertEqualCase(t, rawText, 0, CF996A)
}
