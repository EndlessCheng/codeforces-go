package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1188/B
// https://codeforces.com/problemset/status/1188/problem/B
func TestCF1188B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 0
0 1 2
outputCopy
1
inputCopy
6 7 2
1 2 3 4 5 6
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1188B)
}
