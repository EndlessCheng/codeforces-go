package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/597/C
// https://codeforces.com/problemset/status/597/problem/C
func TestCF597C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
1
2
3
5
4
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF597C)
}
