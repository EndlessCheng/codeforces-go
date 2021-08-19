package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1491/C
// https://codeforces.com/problemset/status/1491/problem/C
func TestCF1491C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7
1 4 2 2 2 2 2
2
2 3
5
1 1 1 1 1
outputCopy
4
3
0`
	testutil.AssertEqualCase(t, rawText, 1, CF1491C)
}
