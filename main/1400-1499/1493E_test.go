package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1493/problem/E
// https://codeforces.com/problemset/status/1493/problem/E
func TestCF1493E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
0010011
1111010
outputCopy
1111111
inputCopy
4
1010
1101
outputCopy
1101`
	testutil.AssertEqualCase(t, rawText, 0, CF1493E)
}
