package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/407/problem/A
// https://codeforces.com/problemset/status/407/problem/A
func TestCF407A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
NO
inputCopy
5 5
outputCopy
YES
2 1
5 5
-2 4
inputCopy
5 10
outputCopy
YES
-10 4
-2 -2
1 2
inputCopy
15 20
outputCopy
YES
0 0
12 9
-12 16`
	testutil.AssertEqualCase(t, rawText, 0, CF407A)
}
