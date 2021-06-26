package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1039/C
// https://codeforces.com/problemset/status/1039/problem/C
func TestCF1039C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 2
0 1 0 1
1 2
2 3
3 4
4 1
outputCopy
50
inputCopy
4 5 3
7 1 7 2
1 2
2 3
3 4
4 1
2 4
outputCopy
96`
	testutil.AssertEqualCase(t, rawText, 0, CF1039C)
}
