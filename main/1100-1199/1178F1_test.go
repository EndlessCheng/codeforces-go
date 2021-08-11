package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1178/F1
// https://codeforces.com/problemset/status/1178/problem/F1
func TestCF1178F1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2 3
outputCopy
5
inputCopy
7 7
4 5 1 6 2 3 7
outputCopy
165`
	testutil.AssertEqualCase(t, rawText, 0, CF1178F1)
}
