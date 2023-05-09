package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/997/A
// https://codeforces.com/problemset/status/997/problem/A
func TestCF997A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 10
01000
outputCopy
11
inputCopy
5 10 1
01000
outputCopy
2
inputCopy
7 2 3
1111111
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF997A)
}
