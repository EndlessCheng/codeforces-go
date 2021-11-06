package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1603/B
// https://codeforces.com/problemset/status/1603/problem/B
func TestCF1603B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 8
4 2
420 420
69420 42068
outputCopy
4
10
420
9969128`
	testutil.AssertEqualCase(t, rawText, 0, CF1603B)
}
