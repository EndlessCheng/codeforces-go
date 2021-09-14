package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1270/C
// https://codeforces.com/problemset/status/1270/problem/C
func TestCF1270C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 2 3 6
1
8
2
1 1
outputCopy
0

2
4 4
3
2 6 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1270C)
}
