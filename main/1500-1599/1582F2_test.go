package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1582/F2
// https://codeforces.com/problemset/status/1582/problem/F2
func TestCF1582F2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 2 2 4
outputCopy
4
0 2 4 6 
inputCopy
8
1 0 1 7 12 5 3 2
outputCopy
12
0 1 2 3 4 5 6 7 10 11 12 13 `
	testutil.AssertEqualCase(t, rawText, 0, CF1582F2)
}
