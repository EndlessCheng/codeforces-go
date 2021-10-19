package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/D2
// https://codeforces.com/problemset/status/1593/problem/D2
func TestCF1593D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6
48 13 22 -15 16 35
8
-1 0 1 -1 0 1 -1 0
4
100 -1000 -1000 -1000
4
1 1 1 1
outputCopy
13
2
-1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1593D2)
}
