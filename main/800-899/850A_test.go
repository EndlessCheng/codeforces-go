package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/850/problem/A
// https://codeforces.com/problemset/status/850/problem/A
func TestCF850A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
0 0 0 0 0
1 0 0 0 0
0 1 0 0 0
0 0 1 0 0
0 0 0 1 0
0 0 0 0 1
outputCopy
1
1
inputCopy
3
0 0 1 2 0
0 0 9 2 0
0 0 5 9 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF850A)
}
