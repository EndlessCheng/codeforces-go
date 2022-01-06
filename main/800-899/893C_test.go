package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/893/problem/C
// https://codeforces.com/problemset/status/893/problem/C
func TestCF893C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
2 5 3 4 8
1 4
4 5
outputCopy
10
inputCopy
10 0
1 2 3 4 5 6 7 8 9 10
outputCopy
55
inputCopy
10 5
1 6 2 7 3 8 4 9 5 10
1 2
3 4
5 6
7 8
9 10
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF893C)
}
