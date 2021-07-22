package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1120/C
// https://codeforces.com/problemset/status/1120/problem/C
func TestCF1120C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 1
aba
outputCopy
7
inputCopy
4 1 1
abcd
outputCopy
4
inputCopy
4 10 1
aaaa
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, 0, CF1120C)
}
