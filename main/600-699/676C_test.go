package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/676/C
// https://codeforces.com/problemset/status/676/problem/C
func TestCF676C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
abba
outputCopy
4
inputCopy
8 1
aabaabaa
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF676C)
}
