package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/660/C
// https://codeforces.com/problemset/status/660/problem/C
func TestCF660C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 1
1 0 0 1 1 0 1
outputCopy
4
1 0 0 1 1 1 1
inputCopy
10 2
1 0 0 1 0 1 0 1 0 1
outputCopy
5
1 0 0 1 1 1 1 1 0 1`
	testutil.AssertEqualCase(t, rawText, 0, CF660C)
}
