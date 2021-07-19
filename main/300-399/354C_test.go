package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/354/C
// https://codeforces.com/problemset/status/354/problem/C
func TestCF354C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 1
3 6 10 12 13 16
outputCopy
3
inputCopy
5 3
8 21 52 15 77
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF354C)
}
