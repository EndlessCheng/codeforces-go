package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/49/D
// https://codeforces.com/problemset/status/49/problem/D
func TestCF49D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
111010
outputCopy
1
inputCopy
5
10001
outputCopy
1
inputCopy
7
1100010
outputCopy
2
inputCopy
5
00100
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF49D)
}
