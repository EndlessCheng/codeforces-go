package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/F
// https://codeforces.com/problemset/status/1593/problem/F
func TestCF1593F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 3 13
02165
4 2 1
1357
8 1 1
12345678
2 7 9
90
outputCopy
RBRBR
-1
BBRRRRBB
BR`
	testutil.AssertEqualCase(t, rawText, 0, CF1593F)
}
