package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1006/D
// https://codeforces.com/problemset/status/1006/problem/D
func TestCF1006D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
abacaba
bacabaa
outputCopy
4
inputCopy
5
zcabd
dbacz
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1006D)
}
