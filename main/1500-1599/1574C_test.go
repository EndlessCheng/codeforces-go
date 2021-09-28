package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1574/C
// https://codeforces.com/problemset/status/1574/problem/C
func TestCF1574C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 6 2 3
5
3 12
7 9
4 14
1 10
8 7
outputCopy
1
2
4
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1574C)
}
