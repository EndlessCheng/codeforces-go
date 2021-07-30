package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1554/problem/D
// https://codeforces.com/problemset/status/1554/problem/D
func TestCF1554D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
5
9
19
outputCopy
abc
diane
bbcaabbba
youarethecutestuwuu
inputCopy
4
1
2
3
4
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1554D)
}
