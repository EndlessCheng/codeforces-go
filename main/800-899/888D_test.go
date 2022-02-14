package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/888/problem/D
// https://codeforces.com/problemset/status/888/problem/D
func TestCF888D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 1
outputCopy
1
inputCopy
4 2
outputCopy
7
inputCopy
5 3
outputCopy
31
inputCopy
5 4
outputCopy
76`
	testutil.AssertEqualCase(t, rawText, 0, CF888D)
}
