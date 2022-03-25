package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1157/problem/B
// https://codeforces.com/problemset/status/1157/problem/B
func TestCF1157B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1337
1 2 5 4 6 6 3 1 9
outputCopy
1557
inputCopy
5
11111
9 8 7 6 5 4 3 2 1
outputCopy
99999
inputCopy
2
33
1 1 1 1 1 1 1 1 1
outputCopy
33`
	testutil.AssertEqualCase(t, rawText, 0, CF1157B)
}
