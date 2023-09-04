package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/960/B
// https://codeforces.com/problemset/status/960/problem/B
func TestCF960B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 0 0
1 2
2 3
outputCopy
2
inputCopy
2 1 0
1 2
2 2
outputCopy
0
inputCopy
2 5 7
3 4
14 4
outputCopy
1
inputCopy
5 5 5
0 0 0 0 0
0 0 0 0 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF960B)
}
