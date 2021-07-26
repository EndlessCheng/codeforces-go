package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/628/D
// https://codeforces.com/problemset/status/628/problem/D
func TestCF628D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 6
10
99
outputCopy
8
inputCopy
2 0
1
9
outputCopy
4
inputCopy
19 7
1000
9999
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 2, CF628D)
}
