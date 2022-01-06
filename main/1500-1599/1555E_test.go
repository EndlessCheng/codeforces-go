package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1555/E
// https://codeforces.com/problemset/status/1555/problem/E
func TestCF1555E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 12
1 5 5
3 4 10
4 10 6
11 12 5
10 12 3
outputCopy
3
inputCopy
1 10
1 10 23
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1555E)
}
