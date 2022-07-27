package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1203/C
// https://codeforces.com/problemset/status/1203/problem/C
func TestCF1203C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
outputCopy
1
inputCopy
6
6 90 12 18 30 18
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1203C)
}
