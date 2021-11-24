package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/496/B
// https://codeforces.com/problemset/status/496/problem/B
func TestCF496B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
579
outputCopy
024
inputCopy
4
2014
outputCopy
0142`
	testutil.AssertEqualCase(t, rawText, 0, CF496B)
}
