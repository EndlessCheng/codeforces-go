package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1811/E
// https://codeforces.com/problemset/status/1811/problem/E
func TestCF1811E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3
5
22
10
100
12345
827264634912
outputCopy
3
6
25
11
121
18937
2932285320890`
	testutil.AssertEqualCase(t, rawText, 0, CF1811E)
}
