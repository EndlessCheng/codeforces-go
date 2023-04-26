package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/915/F
// https://codeforces.com/problemset/status/915/problem/F
func TestCF915F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2 3 1
1 2
1 3
1 4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF915F)
}
