package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/D1
// https://codeforces.com/problemset/status/1593/problem/D1
func TestCF1593D1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
1 5 3 1 1 5
8
-1 0 1 -1 0 1 -1 0
4
100 -1000 -1000 -1000
outputCopy
2
1
1100`
	testutil.AssertEqualCase(t, rawText, 0, CF1593D1)
}
