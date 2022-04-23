package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1672/C
// https://codeforces.com/problemset/status/1672/problem/C
func TestCF1672C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 1 1 1 1
5
2 1 1 1 2
6
1 1 2 3 3 4
6
1 2 1 4 5 4
outputCopy
2
1
2
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1672C)
}
