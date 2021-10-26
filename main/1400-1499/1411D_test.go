package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1411/D
// https://codeforces.com/problemset/status/1411/problem/D
func TestCF1411D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0?1
2 3
outputCopy
4
inputCopy
?????
13 37
outputCopy
0
inputCopy
?10?
239 7
outputCopy
28
inputCopy
01101001
5 7
outputCopy
96`
	testutil.AssertEqualCase(t, rawText, 0, CF1411D)
}
