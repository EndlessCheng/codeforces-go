package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/385/C
// https://codeforces.com/problemset/status/385/problem/C
func TestCF385C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 5 7 10 14 15
3
2 11
3 12
4 4
outputCopy
9
7
0
inputCopy
7
2 3 5 7 11 4 8
2
8 10
2 123
outputCopy
0
7`
	testutil.AssertEqualCase(t, rawText, 0, CF385C)
}
