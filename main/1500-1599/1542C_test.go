package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1542/C
// https://codeforces.com/problemset/status/1542/problem/C
func TestCF1542C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1
2
3
4
10
10000000000000000
outputCopy
2
5
7
10
26
366580019`
	testutil.AssertEqualCase(t, rawText, 0, CF1542C)
}
