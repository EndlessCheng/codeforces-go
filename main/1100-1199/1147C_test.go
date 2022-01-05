package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1147/C
// https://codeforces.com/problemset/status/1147/problem/C
func TestCF1147C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
8 8
outputCopy
Bob
inputCopy
4
3 1 4 1
outputCopy
Alice`
	testutil.AssertEqualCase(t, rawText, 0, CF1147C)
}
