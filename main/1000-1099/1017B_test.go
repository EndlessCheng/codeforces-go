package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1017/B
// https://codeforces.com/problemset/status/1017/problem/B
func TestCF1017B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
01011
11001
outputCopy
4
inputCopy
6
011000
010011
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1017B)
}
