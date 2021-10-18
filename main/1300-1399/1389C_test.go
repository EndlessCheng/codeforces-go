package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1389/C
// https://codeforces.com/problemset/status/1389/problem/C
func TestCF1389C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
95831
100120013
252525252525
outputCopy
3
5
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1389C)
}
