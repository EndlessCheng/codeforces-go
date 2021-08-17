package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1203/F2
// https://codeforces.com/problemset/status/1203/problem/F2
func TestCF1203F2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
4 6
10 -2
8 -1
outputCopy
3
inputCopy
5 20
45 -6
34 -15
10 34
1 27
40 -45
outputCopy
5
inputCopy
3 2
300 -300
1 299
1 123
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1203F2)
}
