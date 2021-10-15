package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1555/C
// https://codeforces.com/problemset/status/1555/problem/C
func TestCF1555C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 3 7
3 5 1
3
1 3 9
3 5 1
1
4
7
outputCopy
7
8
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1555C)
}
