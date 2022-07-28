package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1213/F
// https://codeforces.com/problemset/status/1213/problem/F
func TestCF1213F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2 3
1 3 2
outputCopy
YES
abb`
	testutil.AssertEqualCase(t, rawText, 0, CF1213F)
}
