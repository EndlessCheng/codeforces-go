package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/962/problem/A
// https://codeforces.com/problemset/status/962/problem/A
func TestCF962A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 3 2 1
outputCopy
2
inputCopy
6
2 2 2 2 2 2
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF962A)
}
