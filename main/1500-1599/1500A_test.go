package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1500/problem/A
// https://codeforces.com/problemset/status/1500/problem/A
func TestCF1500A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 1 5 2 7 4
outputCopy
YES
2 3 1 6 
inputCopy
5
1 3 1 9 20
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1500A)
}
