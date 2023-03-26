package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1535/C
// https://codeforces.com/problemset/status/1535/problem/C
func TestCF1535C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0?10
???
?10??1100
outputCopy
8
6
25`
	testutil.AssertEqualCase(t, rawText, 0, CF1535C)
}
