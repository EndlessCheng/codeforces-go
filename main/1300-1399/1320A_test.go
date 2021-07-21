package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1320/A
// https://codeforces.com/problemset/status/1320/problem/A
func TestCF1320A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
10 7 1 9 10 15
outputCopy
26
inputCopy
1
400000
outputCopy
400000
inputCopy
7
8 9 26 11 12 29 14
outputCopy
55`
	testutil.AssertEqualCase(t, rawText, 0, CF1320A)
}
