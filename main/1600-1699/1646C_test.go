package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1646/C
// https://codeforces.com/problemset/status/1646/problem/C
func TestCF1646C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7
11
240
17179869184
outputCopy
2
3
4
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1646C)
}
