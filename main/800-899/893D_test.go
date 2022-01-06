package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/893/problem/D
// https://codeforces.com/problemset/status/893/problem/D
func TestCF893D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 10
-1 5 0 -5 3
outputCopy
0
inputCopy
3 4
-10 0 20
outputCopy
-1
inputCopy
5 10
-5 0 10 -11 0
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF893D)
}
