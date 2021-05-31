package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/788/A
// https://codeforces.com/problemset/status/788/problem/A
func TestCF788A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 4 2 3 1
outputCopy
3
inputCopy
4
1 5 4 7
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF788A)
}
