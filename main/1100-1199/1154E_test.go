package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1154/E
// https://codeforces.com/problemset/status/1154/problem/E
func TestCF1154E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
2 4 5 3 1
outputCopy
11111
inputCopy
5 1
2 1 3 5 4
outputCopy
22111
inputCopy
7 1
7 2 1 3 5 4 6
outputCopy
1121122
inputCopy
5 1
2 4 5 3 1
outputCopy
21112`
	testutil.AssertEqualCase(t, rawText, -1, CF1154E)
}
