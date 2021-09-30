package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1561/C
// https://codeforces.com/problemset/status/1561/problem/C
func TestCF1561C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1
1 42
2
3 10 15 8
2 12 11
outputCopy
43
13`
	testutil.AssertEqualCase(t, rawText, 0, CF1561C)
}
