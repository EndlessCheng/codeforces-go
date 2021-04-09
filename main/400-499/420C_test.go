package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/420/problem/C
// https://codeforces.com/problemset/status/420/problem/C
func TestCF420C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
2 3
1 4
1 4
2 1
outputCopy
6
inputCopy
8 6
5 6
5 7
5 8
6 2
2 1
7 3
1 3
1 4
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF420C)
}
