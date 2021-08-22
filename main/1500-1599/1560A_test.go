package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1560/A
// https://codeforces.com/problemset/status/1560/problem/A
func TestCF1560A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
1
2
3
4
5
6
7
8
9
1000
outputCopy
1
2
4
5
7
8
10
11
14
1666`
	testutil.AssertEqualCase(t, rawText, 0, CF1560A)
}
