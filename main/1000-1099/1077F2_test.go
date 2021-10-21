package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1077/F2
// https://codeforces.com/problemset/status/1077/problem/F2
func TestCF1077F2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2 3
5 1 3 10 1
outputCopy
18
inputCopy
6 1 5
10 30 30 70 10 10
outputCopy
-1
inputCopy
4 3 1
1 100 1 1
outputCopy
100`
	testutil.AssertEqualCase(t, rawText, 0, CF1077F2)
}
