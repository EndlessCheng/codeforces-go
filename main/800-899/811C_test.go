package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/811/C
// https://codeforces.com/problemset/status/811/problem/C
func TestCF811C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 4 2 5 2 3
outputCopy
14
inputCopy
9
5 1 3 1 5 2 4 2 5
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF811C)
}
