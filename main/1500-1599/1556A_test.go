package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1556/A
// https://codeforces.com/problemset/status/1556/problem/A
func TestCF1556A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2
3 5
5 3
6 6
8 0
0 0
outputCopy
-1
2
2
1
2
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1556A)
}
