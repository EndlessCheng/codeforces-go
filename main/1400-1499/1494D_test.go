package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1494/problem/D
// https://codeforces.com/problemset/status/1494/problem/D
func TestCF1494D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 5 7
5 1 7
7 7 4
outputCopy
5
2 1 4 7 5 
4
1 5
2 5
5 4
3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1494D)
}
