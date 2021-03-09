package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1493/problem/A
// https://codeforces.com/problemset/status/1493/problem/A
func TestCF1493A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 2
5 3
1 1
outputCopy
2
3 1 
3
4 5 2 
0
`
	testutil.AssertEqualCase(t, rawText, 0, CF1493A)
}
