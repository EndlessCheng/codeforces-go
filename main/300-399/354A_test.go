package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/354/A
// https://codeforces.com/problemset/status/354/problem/A
func TestCF354A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 4 19 1
42 3 99
outputCopy
576
inputCopy
4 7 2 3 9
1 2 3 4
outputCopy
34`
	testutil.AssertEqualCase(t, rawText, 0, CF354A)
}
