package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1514/problem/C
// https://codeforces.com/problemset/status/1514/problem/C
func TestCF1514C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
outputCopy
3
1 2 3 
inputCopy
8
outputCopy
4
1 3 5 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1514C)
}
