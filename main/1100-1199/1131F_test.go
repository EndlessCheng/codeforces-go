package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1131/F
// https://codeforces.com/problemset/status/1131/problem/F
func TestCF1131F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 4
2 5
3 1
4 5
outputCopy
3 1 4 2 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1131F)
}
