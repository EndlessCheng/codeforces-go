package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/711/D
// https://codeforces.com/problemset/status/711/problem/D
func TestCF711D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 1
outputCopy
6
inputCopy
4
2 1 1 1
outputCopy
8
inputCopy
5
2 4 2 5 3
outputCopy
28`
	testutil.AssertEqualCase(t, rawText, 0, CF711D)
}
