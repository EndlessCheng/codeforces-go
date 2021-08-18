package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1557/C
// https://codeforces.com/problemset/status/1557/problem/C
func TestCF1557C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 1
2 1
4 0
outputCopy
5
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1557C)
}
