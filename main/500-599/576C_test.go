package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/576/C
// https://codeforces.com/problemset/status/576/problem/C
func TestCF576C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 7
8 10
3 4
5 0
9 12
outputCopy
4 3 1 2 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF576C)
}
