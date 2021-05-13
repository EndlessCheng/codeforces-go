package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/715/B
// https://codeforces.com/problemset/status/715/problem/B
func TestCF715B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 13 0 4
0 1 5
2 1 2
3 2 3
1 4 0
4 3 4
outputCopy
YES
0 1 5
2 1 2
3 2 3
1 4 8
4 3 4
inputCopy
2 1 123456789 0 1
0 1 0
outputCopy
YES
0 1 123456789
inputCopy
2 1 999999999 1 0
0 1 1000000000
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF715B)
}
