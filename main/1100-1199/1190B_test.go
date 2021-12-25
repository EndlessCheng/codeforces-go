package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1190/B
// https://codeforces.com/problemset/status/1190/problem/B
func TestCF1190B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
0
outputCopy
cslnb
inputCopy
2
1 0
outputCopy
cslnb
inputCopy
2
2 2
outputCopy
sjfnb
inputCopy
3
2 3 1
outputCopy
sjfnb
inputCopy
3
0 0 6
outputCopy
cslnb
inputCopy
3
4 5 5
outputCopy
cslnb`
	testutil.AssertEqualCase(t, rawText, 0, CF1190B)
}
