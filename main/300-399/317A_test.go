package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/317/A
// https://codeforces.com/problemset/status/317/problem/A
func TestCF317A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 2 5
outputCopy
2
inputCopy
-1 4 15
outputCopy
4
inputCopy
0 -1 5
outputCopy
-1
inputCopy
999999999 -1000000000 1000000000
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF317A)
}
