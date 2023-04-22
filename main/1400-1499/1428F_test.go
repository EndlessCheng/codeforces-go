package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1428/F
// https://codeforces.com/problemset/status/1428/problem/F
func TestCF1428F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0110
outputCopy
12
inputCopy
7
1101001
outputCopy
30
inputCopy
12
011100011100
outputCopy
156`
	testutil.AssertEqualCase(t, rawText, 0, CF1428F)
}
