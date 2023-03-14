package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/687/C
// https://codeforces.com/problemset/status/687/problem/C
func TestCF687C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 18
5 6 1 10 12 2
outputCopy
16
0 1 2 3 5 6 7 8 10 11 12 13 15 16 17 18 
inputCopy
3 50
25 25 50
outputCopy
3
0 25 50 `
	testutil.AssertEqualCase(t, rawText, 0, CF687C)
}
