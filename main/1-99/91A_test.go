package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/91/A
// https://codeforces.com/problemset/status/91/problem/A
func TestCF91A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abc
xyz
outputCopy
-1
inputCopy
abcd
dabc
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF91A)
}
