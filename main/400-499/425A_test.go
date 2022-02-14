package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/425/problem/A
// https://codeforces.com/problemset/status/425/problem/A
func TestCF425A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 2
10 -1 2 2 2 2 2 2 -1 10
outputCopy
32
inputCopy
5 10
-1 -1 -1 -1 -1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF425A)
}
